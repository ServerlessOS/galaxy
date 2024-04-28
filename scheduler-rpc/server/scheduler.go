package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServerlessOS/galaxy/constant"
	pb "github.com/ServerlessOS/galaxy/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"math/rand/v2"
	"net"
	"os"
	"scheduler_rpc/internal"
	"scheduler_rpc/internal/cache"
	"strconv"
	"strings"
	"time"
)

var (
	nodeView                *cache.Cache
	nodeResourceUpdateQueue *cache.FIFO
	RequestQueue            *internal.PriorityQueue
	NotDeployed             map[int64]struct{}
	ConnCache               *LRUCache
	PeerSchedulers          *internal.SchedulerView

	schedulerName string
	localRpcAddr  string
	gatewayAddr   string
	localIp       = getLocalIPv4().String()
)

type SchedulerServer struct{}

func init() {
	RequestQueue = internal.NewPriorityQueue()
	nodeResourceUpdateQueue = cache.NewFIFO()
	nodeView = cache.NewCache()
	NotDeployed = make(map[int64]struct{})
	ConnCache = NewLRUCache(20)
	PeerSchedulers = internal.NewSchedulerView()
	Cmd.Flags().StringVarP(&localRpcAddr, "localRpcAddr", "r", "0.0.0.0:"+constant.SchedulerPort, "The addr used for binding to the RPC server. ")
	Cmd.Flags().StringVarP(&gatewayAddr, "gatewayAddr", "g", "", "The address information of the gateway needs to be registered with the gateway to work properly. ")

}

var Cmd = &cobra.Command{
	Use:   "scheduler",
	Short: `初始化scheduler程序`,
	//本函数用于执行命令并返回错误
	RunE: func(cmd *cobra.Command, args []string) error {
		schedulerName = strconv.Itoa(int(rand.Uint32()))
		var errChanRpc chan error
		if !cmd.Flags().Changed("gatewayAddr") {
			return errors.New("gatewayAddr is required")
		}
		go rpcServer(errChanRpc)
		register()
		err := <-errChanRpc
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
			return err
		}
		return nil
	},
}

// Run 提供给顶层用于启动cobra根命令
func Run(cmd *cobra.Command) (code int) {
	err := cmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return 0
}
func register() {
	//通过gateway向顶层控制器注册
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	connGateway, err := grpc.Dial(gatewayAddr+":"+constant.GatewayRpcPort, grpc.WithInsecure(), grpc.WithTimeout(time.Second*3))
	if err != nil {
		log.Fatalln("dial gateway error:", err)
	}
	client := pb.NewGatewayClient(connGateway)
	if isIPAddress(localRpcAddr) {
		tcpAddr, err := net.ResolveTCPAddr("tcp", localRpcAddr)
		_, err = client.Register(ctx, &pb.RegisterReq{
			Type:    3, //    coordinator = 0; funcManager = 1;
			Name:    schedulerName,
			Address: tcpAddr.IP.String(),
		})
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		_, err = client.Register(ctx, &pb.RegisterReq{
			Type:    3, //    coordinator = 0; funcManager = 1;
			Name:    schedulerName,
			Address: localIp,
		})
		if err != nil {
			log.Fatalln(err)
		}
	}

}
func rpcServer(errChannel chan<- error) {
	lis, err := net.Listen("tcp", localRpcAddr)
	if err != nil {
		errChannel <- err
	}
	// 实例化grpc服务端
	s := grpc.NewServer()

	// 在gRPC服务器注册服务
	pb.RegisterSchedulerServer(s, &SchedulerServer{})

	// 启动grpc服务
	err = s.Serve(lis)
	errChannel <- err
}
func (s SchedulerServer) PeerSchedulerUpdate(ctx context.Context, update *pb.PeerSchedulersUpdate) (*pb.PeerSchedulersUpdateReply, error) {

	for _, peerScheduler := range update.List {
		sitem := &internal.SchedulerInfo{
			NodeName: peerScheduler.NodeName,
			Address:  peerScheduler.Address,
			MemLimit: -1,
		}
		if update.Action == "ADD" {
			fmt.Printf("Add new peer scheduler %s\n", sitem.Address)
			PeerSchedulers.Add(sitem)
		} else {
			PeerSchedulers.Delete(sitem)
		}
	}
	return &pb.PeerSchedulersUpdateReply{State: 0}, nil
}

func (s SchedulerServer) PeerSchedule(ctx context.Context, request *pb.ScheduleRequest) (*pb.ScheduleReply, error) {

	nodeView.Lock.Lock()
	defer nodeView.Lock.Unlock()
	for nodeName, node := range nodeView.Cache {
		if node.HaveCpu >= request.RequireCpu && node.HaveMem >= request.RequireMem {
			log.Printf("deploy -%d-%s to node %s at -%d\n", request.RequestId, request.FuncName, nodeName, time.Now().UnixNano()/1e6)
			node.HaveMem -= request.RequireMem
			node.HaveCpu -= request.RequireCpu
			nodeView.Set(nodeName, node)
			// deploy to node
			conn := ConnCache.Get(node.Address)
			if conn != nil {
				fmt.Printf("Cache hitted!\n")
			} else {
				conn, _ = grpc.Dial(fmt.Sprintf("%s:16446", node.Address), grpc.WithInsecure())
				ConnCache.Put(node.Address, conn)
				fmt.Printf("Cache miss!\n")
			}
			clinet := pb.NewNodeClient(conn)
			_, _ = clinet.Deploy(context.Background(), &pb.InstanceDeploy{List: []*pb.NodeInstanceInfo{
				{
					RequestId:      request.RequestId,
					FuncName:       request.FuncName,
					DispatcherAddr: request.DispatcherAddr,
				},
			}})
			return &pb.ScheduleReply{
				RequestId:      request.RequestId,
				FuncName:       request.FuncName,
				DeployPosition: node.Address,
			}, nil
		}
	}

	return &pb.ScheduleReply{
		RequestId:      request.RequestId,
		FuncName:       request.FuncName,
		DeployPosition: "",
	}, nil
}

func (s SchedulerServer) Schedule(ctx context.Context, requests *pb.ScheduleRequestList) (*pb.ScheduleReply, error) {
	// 根据将请求入列
	priority := time.Now().UnixNano() / 1e6
	for _, request := range requests.List {
		log.Printf("Scheduler receive -%d- at -%d\n", request.RequestId, priority)
		//fmt.Printf("Scheduler receive -%d- at -%d\n", request.RequestId, priority)
		reqestInfo := &internal.RequestInfo{RequestId: request.RequestId, FunctionName: request.FuncName, RequireCpu: request.RequireCpu, RequireMem: request.RequireMem, DispatcherAddr: request.DispatcherAddr}
		if strings.Contains(request.FuncName, "galaxy-") {
			priority = 0
		}
		//fmt.Printf("Receive new %d:%s\n", request.RequestId, request.FuncName)
		RequestQueue.Push(&internal.RequestItem{reqestInfo, priority})
	}

	return &pb.ScheduleReply{
		RequestId:      0,
		FuncName:       "",
		DeployPosition: "",
	}, nil
}

func (s SchedulerServer) UpadateNodeResource(ctx context.Context, update *pb.NodeResourceUpdate) (*pb.NodeResourceReply, error) {
	nodeView.Lock.Lock()
	action := update.Action
	if action == "ADD" {
		//fmt.Printf("Add %d nodes\n", len(update.List))
		// add the new node to the scheduler list
		for _, nodeResource := range update.List {

			node := &internal.NodeResource{
				NodeName: nodeResource.NodeName,
				HaveCpu:  nodeResource.HaveCpu,
				HaveMem:  nodeResource.HaveMem,
				Address:  nodeResource.Address,
				Port:     nodeResource.Port,
			}
			nodeView.Set(nodeResource.NodeName, node)
		}
	} else if action == "DELETE" {
		for i := 0; i < len(update.List); i++ {
			n := nodeView.Get(update.List[i].NodeName)
			update.List[i].HaveCpu = n.HaveCpu
			update.List[i].HaveMem = n.HaveMem
			nodeView.Delete(update.List[i].NodeName)
		}
	}
	nodeView.Lock.Unlock()
	if action == "DELETE" {
		TransResource(update)
	}
	return &pb.NodeResourceReply{
		State: 0,
	}, nil
}
func TransResource(update *pb.NodeResourceUpdate) {

	// 获取当前时间
	currentTime := time.Now()
	// 转换为毫秒级时间戳
	milliseconds := currentTime.UnixNano() / int64(time.Millisecond)
	fmt.Printf("Move %d2 nodes to %s at %d\n", len(update.List), update.TargetAddr, milliseconds)
	// 输出毫秒级时间戳

	conn, _ := grpc.Dial(fmt.Sprintf("%s:16445", update.TargetAddr), grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewSchedulerClient(conn)

	update.Action = "ADD"
	resp, _ := client.UpadateNodeResource(context.Background(), update)
	log.Printf("client.UpadateNodeResource resp: %d", resp.State)
}
func Schedule() {

	fmt.Printf("Begin to schedule...\n")
	for {
		// 拿到请求
		i := RequestQueue.Pop()
		it := i.Value
		hasDeployed := false
		nodeView.Lock.Lock()
		for nodeName, node := range nodeView.Cache {
			if node.HaveCpu >= it.RequireCpu && node.HaveMem >= it.RequireMem {
				log.Printf("deploy -%d-%s to node %s at -%d\n", it.RequestId, it.FunctionName, nodeName, time.Now().UnixNano()/1e6)
				node.HaveMem -= it.RequireMem
				node.HaveCpu -= it.RequireCpu
				nodeView.Set(nodeName, node)
				// deploy to node
				conn := ConnCache.Get(node.Address)
				if conn != nil {
					fmt.Printf("Cache hitted!\n")
				} else {
					conn, _ = grpc.Dial(fmt.Sprintf("%s:16446", node.Address), grpc.WithInsecure())
					ConnCache.Put(node.Address, conn)
					fmt.Printf("Cache miss!\n")
				}
				clinet := pb.NewNodeClient(conn)
				_, _ = clinet.Deploy(context.Background(), &pb.InstanceDeploy{List: []*pb.NodeInstanceInfo{
					{
						RequestId:      it.RequestId,
						FuncName:       it.FunctionName,
						DispatcherAddr: it.DispatcherAddr,
					},
				}})
				hasDeployed = true
				break
			}
		}
		nodeView.Lock.Unlock()
		if !hasDeployed {
			go ForwardPeerScheudle(it)
		}

	}
}
func ForwardPeerScheudle(request *internal.RequestInfo) {
	peerScheduler1 := PeerSchedulers.GetSchedulerAddr(request.RequireMem)
	if peerScheduler1 != "" {
		conn, _ := grpc.Dial(fmt.Sprintf("%s:16445", peerScheduler1), grpc.WithInsecure())
		client := pb.NewSchedulerClient(conn)
		r := &pb.ScheduleRequest{
			RequestId:      request.RequestId,
			FuncName:       request.FunctionName,
			RequireCpu:     request.RequireCpu,
			RequireMem:     request.RequireMem,
			DispatcherAddr: request.DispatcherAddr,
		}
		peerReply, _ := client.PeerSchedule(context.Background(), r)
		if peerReply.DeployPosition != "" {
			fmt.Printf("Request %d success to probe\n", peerReply.RequestId)
			return
		} else {
			PeerSchedulers.SetSchedulerLimit(peerScheduler1, r.RequireMem)
		}
	} else {
		fmt.Printf("Request %d cant find the position to deploy\n", request.RequestId)
		return
	}

	peerScheduler2 := PeerSchedulers.GetSchedulerAddr(request.RequireMem)
	if peerScheduler1 != "" && peerScheduler2 != peerScheduler1 {
		conn, _ := grpc.Dial(fmt.Sprintf("%s:16445", peerScheduler2), grpc.WithInsecure())
		client := pb.NewSchedulerClient(conn)
		r := &pb.ScheduleRequest{
			RequestId:      request.RequestId,
			FuncName:       request.FunctionName,
			RequireCpu:     request.RequireCpu,
			RequireMem:     request.RequireMem,
			DispatcherAddr: request.DispatcherAddr,
		}
		peerReply, _ := client.PeerSchedule(context.Background(), r)
		if peerReply.DeployPosition != "" {
			fmt.Printf("Request %d success to probe\n", peerReply.RequestId)
			return
		} else {
			PeerSchedulers.SetSchedulerLimit(peerScheduler1, r.RequireMem)
		}
	} else {
		fmt.Printf("Scheduler cant find the deployed position for %d:%s.\n", request.RequestId, request.FunctionName)
		//fmt.Printf("Request %d  cant find the other peerScheduler\n", request.RequestId)
	}

}
func isIPAddress(addr string) bool {
	ip, _, err := net.SplitHostPort(addr)
	if err != nil {
		return false
	}
	return net.ParseIP(ip) != nil
}
func getLocalIPv4() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println(ipNet.IP.String())
				return ipNet.IP
			}
		}
	}
	return nil
}
