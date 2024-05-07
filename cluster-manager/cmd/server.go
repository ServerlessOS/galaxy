package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServerlessOS/galaxy/constant"
	"github.com/ServerlessOS/galaxy/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"math/rand/v2"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	localRpcAddr       string
	gatewayAddr        string
	clusterManagerName string
	ExitCh             = make(chan int)
	localIp            = getLocalIPv4().String()
)

type rpcServerProcess struct{}

var Cmd = &cobra.Command{
	Use:   "clusterManager",
	Short: `初始化clusterManager程序`,
	//本函数用于执行命令并返回错误
	RunE: func(cmd *cobra.Command, args []string) error {
		clusterManagerName = strconv.Itoa(int(rand.Uint32()))
		var errChanRpc chan error
		if !cmd.Flags().Changed("gatewayAddr") {
			return errors.New("clusterManagerAddr is required")
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
	Cmd.Flags().StringVarP(&localRpcAddr, "localRpcAddr", "r", "0.0.0.0:"+constant.FuncManagerPort, "The addr used for binding to the RPC server. ")
	Cmd.Flags().StringVarP(&gatewayAddr, "gatewayAddr", "g", "", "The address information of the gateway needs to be registered with the gateway to work properly. ")
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
	log.Println("Init clusterManager success.")
}
func register() {
	//通过gateway向顶层控制器注册
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	connGateway, err := grpc.Dial(gatewayAddr+":"+constant.GatewayRpcPort, grpc.WithInsecure(), grpc.WithTimeout(time.Second*3))
	if err != nil {
		log.Fatalln("dial gateway error:", err)
	}
	client := proto.NewGatewayClient(connGateway)
	tcpAddr, err := net.ResolveTCPAddr("tcp", localRpcAddr)
	_, err = client.Register(ctx, &proto.RegisterReq{
		Type:    1, //    coordinator = 0; funcManager = 1;
		Name:    clusterManagerName,
		Address: tcpAddr.IP.String(),
	})
}

func rpcServer(errChannel chan<- error) {
	lis, err := net.Listen("tcp", localRpcAddr)
	if err != nil {
		errChannel <- err
	}
	// 实例化grpc服务端
	s := grpc.NewServer()

	// 在gRPC服务器注册服务
	proto.RegisterClusterManagerServer(s, &rpcServerProcess{})

	// 启动grpc服务
	err = s.Serve(lis)
	errChannel <- err
}
func (r rpcServerProcess) MoniterUpload(ctx context.Context, req *proto.MoniterUploadReq) (*proto.MoniterUploadResp, error) {
	//TODO implement me
	panic("implement me")
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
