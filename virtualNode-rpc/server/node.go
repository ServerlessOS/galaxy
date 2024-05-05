package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServerlessOS/galaxy/constant"
	pb "github.com/ServerlessOS/galaxy/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"
	"virtualNode_rpc/client"
	"virtualNode_rpc/internal"
)

var (
	DeployQueue = internal.NewFIFO()
	ConnCache   *LRUCache

	nodeName     string
	localRpcAddr string
	gatewayAddr  string
	localIp      = getLocalIPv4().String()
)

func init() {
	ConnCache = NewLRUCache(20)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
	Cmd.Flags().StringVarP(&localRpcAddr, "localRpcAddr", "r", "0.0.0.0:"+constant.NodePort, "The addr used for binding to the RPC server. ")
	Cmd.Flags().StringVarP(&gatewayAddr, "gatewayAddr", "g", "", "The address information of the gateway needs to be registered with the gateway to work properly. ")
}

type NodeServer struct{}

var Cmd = &cobra.Command{
	Use:   "node",
	Short: `初始化node程序`,
	//本函数用于执行命令并返回错误
	RunE: func(cmd *cobra.Command, args []string) error {
		nodeName = strconv.Itoa(int(rand.Uint32()))
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
func register() {
	//通过gateway向顶层控制器注册
	client.DialGatewayClient("default", gatewayAddr)
	tcpAddr, _ := net.ResolveTCPAddr("tcp", localRpcAddr)
	_, err := client.GetGatewayClient().Register(context.Background(), &pb.RegisterReq{
		Type:    4, //    coordinator = 0; funcManager = 1;
		Name:    nodeName,
		Address: tcpAddr.IP.String(),
	})
	if err != nil {
		log.Fatalln(err)
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
	pb.RegisterNodeServer(s, &NodeServer{})

	// 启动grpc服务
	err = s.Serve(lis)
	errChannel <- err
}
func (n NodeServer) Deploy(ctx context.Context, deploy *pb.InstanceDeploy) (*pb.InstanceDeployReply, error) {
	list := deploy.List
	for _, instanceInfo := range list {
		log.Println("deploy instance")
		//resp, _ := client.GetGatewayClient().GetFuncInfo(context.Background(), &pb.GetFuncInfoReq{FuncName: instanceInfo.FuncName}) //此处可以拿到document，但是暂时没做json的解析和配置，所以不使用，仅形式上应用
		//document := resp.FuncInfo
		dockerStart(instanceInfo.FuncName, "") //这里可以填入document，但是需要先预定义document相关的内容
		newInstance := &internal.InstanceInfo{
			RequestId:         instanceInfo.RequestId,
			FuncName:          instanceInfo.FuncName,
			Address:           localIp,
			DispatcherAddress: instanceInfo.DispatcherAddr,
		}
		DeployQueue.Enqueue(newInstance)
		log.Printf("Deploy instace -%d-%s at -%d\n", instanceInfo.RequestId, instanceInfo.FuncName, time.Now().UnixNano()/1e6)
		//fmt.Printf("Deploy instace -%d-%s at -%d\n", instanceInfo.RequestId, instanceInfo.FuncName, time.Now().UnixNano()/1e6)
	}

	return &pb.InstanceDeployReply{State: 0}, nil
}
func dockerStart(name, document string) string {
	//document暂时用不到，如果需要支持例如docker启动选项等内容可以用上
	cmd := exec.Command("docker", "run", name) //todo：权且将name当作imageName，后期最好使用document定义
	out, err := cmd.CombinedOutput()           //todo：注意这是阻塞的
	if err != nil {
		log.Errorln(err)
	}
	log.Println("Command output:", string(out))
	return string(out)
}
func InstanceInfoInform() {
	for {
		instanceInfo := DeployQueue.Dequeue()
		dispatcherAddr := instanceInfo.DispatcherAddress
		// dispatcher 端client调用
		ctx := context.Background()
		conn := ConnCache.Get(dispatcherAddr)
		if conn != nil {
			fmt.Printf("Cache hitted!\n")
		} else {
			conn, _ = grpc.Dial(fmt.Sprintf("%s:16444", dispatcherAddr), grpc.WithInsecure())
			ConnCache.Put(dispatcherAddr, conn)
			fmt.Printf("Cache miss!\n")
		}
		client := pb.NewDispatcherClient(conn)
		_, err := client.UpdateInstanceView(ctx, &pb.InstanceUpdate{
			List: []*pb.InstanceInfo{
				&pb.InstanceInfo{
					RequestId: instanceInfo.RequestId,
					FuncName:  instanceInfo.FuncName,
					Address:   instanceInfo.Address,
				},
			},
			Action: "ADD",
		})
		if err != nil {
			return
		}
		log.Printf("Node update instance -%d- at -%d\n", instanceInfo.RequestId, time.Now().UnixNano()/1e6)
	}
}
func randomIP() net.IP {
	// 生成随机的IPv4地址
	ip := make(net.IP, 4)
	rand.Seed(time.Now().UnixNano())
	for i := range ip {
		ip[i] = byte(rand.Intn(256))
	}
	return ip
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
