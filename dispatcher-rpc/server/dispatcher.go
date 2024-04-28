package server

import (
	"context"
	"dispatcher_rpc/internal"
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
	"strconv"
	"time"
)

var (
	funcView      = internal.NewFuncView() // 支持查找变更 func
	SchedulerView = internal.NewSchedulerView()
	localIp       = getLocalIPv4().String()
	connQueue     = NewConnQueue()
	firstInit     = true

	dispatcherName string
	localRpcAddr   string
	gatewayAddr    string
)

type DispatcherServer struct{}

var Cmd = &cobra.Command{
	Use:   "dispatcher",
	Short: `初始化dispatcher程序`,
	//本函数用于执行命令并返回错误
	RunE: func(cmd *cobra.Command, args []string) error {
		dispatcherName = strconv.Itoa(int(rand.Uint32()))
		var errChanRpc chan error
		if !cmd.Flags().Changed("gatewayAddr") {
			return errors.New("gatewayAddr is required")
		}
		register()
		rpcServer(errChanRpc)
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
func init() {
	Cmd.Flags().StringVarP(&localRpcAddr, "localRpcAddr", "r", ":"+constant.DispatcherPort, "The addr used for binding to the RPC server. ")
	Cmd.Flags().StringVarP(&gatewayAddr, "gatewayAddr", "g", "", "The address information of the gateway needs to be registered with the gateway to work properly. ")
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
			Type:    2, //    coordinator = 0; funcManager = 1;
			Name:    dispatcherName,
			Address: tcpAddr.IP.String(),
		})
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		_, err = client.Register(ctx, &pb.RegisterReq{
			Type:    2, //    coordinator = 0; funcManager = 1;
			Name:    dispatcherName,
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
	pb.RegisterDispatcherServer(s, &DispatcherServer{})

	// 启动grpc服务
	err = s.Serve(lis)
	errChannel <- err
}
func (d DispatcherServer) Statis(ctx context.Context, request *pb.UserRequest) (*pb.UserRequestReply, error) {
	return &pb.UserRequestReply{
		RequestId:   1,
		FuncName:    "request.FuncName",
		Destination: "result",
	}, nil
}

func (d DispatcherServer) UpdateInstanceView(ctx context.Context, update *pb.InstanceUpdate) (*pb.InstanceUpdateReply, error) {
	list := update.List
	for _, v := range list {
		t := time.Now().UnixNano() / 1e6
		item := &internal.FuncInfo{FuncName: v.FuncName,
			Address:   v.Address,
			Timestamp: t,
			State:     true}
		if update.Action == "ADD" {
			funcView.Add(item)
			log.Printf("Add new instance -%d-%s-%s at -%d\n", v.RequestId, item.FuncName, item.Address, time.Now().UnixNano()/1e6)
			//fmt.Printf("Add new instance -%d-%s-%s at -%d\n", v.RequestId, item.FuncName, item.Address, time.Now().UnixNano()/1e6)
		} else if update.Action == "DELETE" {
			funcView.Delete(item)
		}
	}
	return &pb.InstanceUpdateReply{
		State: 0,
	}, nil
}
func randonInitConnection() {
	for i := 0; i < 20; i++ {
		saddr := SchedulerView.GetSchedulerAddr()
		conn, _ := grpc.Dial(fmt.Sprintf("%s:16445", saddr), grpc.WithInsecure())
		connQueue.Enqueue(conn)
	}
}
func addConn(n int) {
	for i := 0; i < n; i++ {
		saddr := SchedulerView.GetSchedulerAddr()
		conn, _ := grpc.Dial(fmt.Sprintf("%s:16445", saddr), grpc.WithInsecure())
		connQueue.Enqueue(conn)
	}
}
func (d DispatcherServer) Dispatch(ctx context.Context, userRequests *pb.UserRequestList) (*pb.UserRequestReply, error) {
	t := time.Now().UnixNano() / 1e6
	per_scheduler := len(userRequests.List) / SchedulerView.GetLen()
	per_scheduler = per_scheduler / 2
	schedulerRequests := make([]*pb.ScheduleRequest, 0)
	for _, request := range userRequests.List {
		result := funcView.Dispatch(request.FuncName)
		if result == "" {
			log.Printf("Need to scale up -%d-%s at -%d\n", request.RequestId, request.FuncName, t)
			//lock.Lock()
			//requestIdList = append(requestIdList, request.RequestId)
			//lock.Unlock()
			schedulerRequests = append(schedulerRequests, &pb.ScheduleRequest{
				RequestId:      request.RequestId,
				FuncName:       request.FuncName,
				RequireCpu:     request.RequireCpu,
				RequireMem:     request.RequireMem,
				DispatcherAddr: localIp,
			})
			//fmt.Printf("len:%d\n", len(schedulerRequests))
			if len(schedulerRequests) >= per_scheduler {
				c := connQueue.Dequeue()
				client := pb.NewSchedulerClient(c)
				//fmt.Printf("Route %d requests to %s\n",len(schedulerRequests),c.GetMethodConfig().)
				_, _ = client.Schedule(context.Background(), &pb.ScheduleRequestList{List: schedulerRequests})
				go addConn(2)
				schedulerRequests = make([]*pb.ScheduleRequest, 0)
			}
		}
	}
	if len(schedulerRequests) > 0 {
		c := connQueue.Dequeue()
		client := pb.NewSchedulerClient(c)
		_, _ = client.Schedule(context.Background(), &pb.ScheduleRequestList{List: schedulerRequests})
		go addConn(2)
		//lock.Lock()
		//fmt.Printf("cnt: %d\n", uniqueKeyCount(requestIdList))
		//lock.Unlock()
	}

	return &pb.UserRequestReply{
		RequestId:   0,
		FuncName:    "0",
		Destination: "0",
	}, nil
}
func (d DispatcherServer) UpdateSchedulerView(ctx context.Context, update *pb.SchedulerViewUpdate) (*pb.SchedulerViewUpdateReply, error) {
	//fmt.Printf("Add new scheduelr \n")
	list := update.List
	for _, v := range list {
		item := &internal.SchedulerInfo{NodeName: v.NodeName, Address: v.Address}
		if update.Action == "ADD" {
			fmt.Printf("Add new scheduelr %s:%s\n", item.NodeName, item.Address)
			SchedulerView.Add(item)

		} else if update.Action == "DELETE" {
			SchedulerView.Delete(item)
		}
	}
	if firstInit {
		go addConn(20)
		firstInit = false
	}
	return &pb.SchedulerViewUpdateReply{State: 0}, nil
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
func uniqueKeyCount(arr []int64) int {
	uniqueKeys := make(map[int64]struct{})

	for _, item := range arr {
		uniqueKeys[item] = struct{}{}
	}

	return len(uniqueKeys)
}
func isIPAddress(addr string) bool {
	ip, _, err := net.SplitHostPort(addr)
	if err != nil {
		return false
	}
	return net.ParseIP(ip) != nil
}
