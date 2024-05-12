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
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	localRpcAddr       string
	gatewayAddr        string
	clusterManagerName string
	ExitCh             = make(chan int)
	localIp            = getLocalIPv4().String()
	gatewayMonitor     = make(map[string]load) //name与负载的映射
	funcManagerMonitor = make(map[string]load)
	dispatcherMonitor  = make(map[string]load)
	schedulerMonitor   = make(map[string]load)
	nodeMonitor        = make(map[string]load)
)

const loadFactorCpu = 0.8
const loadFactorMem = 0.5

type load struct {
	cpuload int
	memload int
}
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
		go patrol()
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
	//gateway = 0;
	//funcManager = 1;
	//dispatcher  = 2;
	//scheduler  = 3;
	//node  = 4;
	switch req.Type {
	case 0:
		name, performance := req.Name, req.Performance
		//performance格式cpuload=?&memload=？
		parts := strings.Split(performance, "&")
		cpuload, _ := strconv.Atoi(parts[0])
		memload, _ := strconv.Atoi(parts[1])
		gatewayMonitor[name] = load{
			cpuload: cpuload,
			memload: memload,
		}
		return &proto.MoniterUploadResp{StatusCode: 0}, nil

	case 1:
		name, performance := req.Name, req.Performance
		//performance格式cpuload=?&memload=？
		parts := strings.Split(performance, "&")
		cpuload, _ := strconv.Atoi(parts[0])
		memload, _ := strconv.Atoi(parts[1])
		funcManagerMonitor[name] = load{
			cpuload: cpuload,
			memload: memload,
		}
		return &proto.MoniterUploadResp{StatusCode: 0}, nil

	case 2:
		name, performance := req.Name, req.Performance
		//performance格式cpuload=?&memload=？
		parts := strings.Split(performance, "&")
		cpuload, _ := strconv.Atoi(parts[0])
		memload, _ := strconv.Atoi(parts[1])
		dispatcherMonitor[name] = load{
			cpuload: cpuload,
			memload: memload,
		}
		return &proto.MoniterUploadResp{StatusCode: 0}, nil

	case 3:
		name, performance := req.Name, req.Performance
		//performance格式cpuload=?&memload=？
		parts := strings.Split(performance, "&")
		cpuload, _ := strconv.Atoi(parts[0])
		memload, _ := strconv.Atoi(parts[1])
		schedulerMonitor[name] = load{
			cpuload: cpuload,
			memload: memload,
		}
		return &proto.MoniterUploadResp{StatusCode: 0}, nil

	case 4:
		name, performance := req.Name, req.Performance
		//performance格式cpuload=?&memload=？
		parts := strings.Split(performance, "&")
		cpuload, _ := strconv.Atoi(parts[0])
		memload, _ := strconv.Atoi(parts[1])
		nodeMonitor[name] = load{
			cpuload: cpuload,
			memload: memload,
		}
		return &proto.MoniterUploadResp{StatusCode: 0}, nil

	default:
		return nil, fmt.Errorf("unknown type")
	}
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

func patrol() {
	for {
		time.Sleep(time.Second * 5)
		var tempCountCpu, tempCountMem int
		for _, v := range gatewayMonitor { //todo：面对实时修改的map使用range可能会出问题
			tempCountCpu += v.cpuload
			tempCountMem += v.memload
		}
		if tempCountCpu*10 > len(gatewayMonitor)*(loadFactorCpu*10) {
			// 发起HTTP GET请求
			resp, err := http.Get(gatewayAddr + ":" + constant.GatewayHttpPort + "/create?" + "funcName=gateway&requireCpu=4&requireMem=4")
			if err != nil {
				fmt.Errorf("扩容失败:", err)
			}
			resp.Body.Close()
		}
		for _, v := range funcManagerMonitor { //todo：面对实时修改的map使用range可能会出问题
			tempCountCpu += v.cpuload
			tempCountMem += v.memload
		}
		if tempCountCpu*10 > len(funcManagerMonitor)*(loadFactorCpu*10) {
			// 发起HTTP GET请求
			resp, err := http.Get(gatewayAddr + ":" + constant.GatewayHttpPort + "/create?" + "funcName=funcManager&requireCpu=4&requireMem=4")
			if err != nil {
				fmt.Errorf("扩容失败:", err)
			}
			resp.Body.Close()
		}
		for _, v := range dispatcherMonitor { //todo：面对实时修改的map使用range可能会出问题
			tempCountCpu += v.cpuload
			tempCountMem += v.memload
		}
		if tempCountCpu*10 > len(dispatcherMonitor)*(loadFactorCpu*10) {
			// 发起HTTP GET请求
			resp, err := http.Get(gatewayAddr + ":" + constant.GatewayHttpPort + "/create?" + "funcName=dispatcher&requireCpu=4&requireMem=4")
			if err != nil {
				fmt.Errorf("扩容失败:", err)
			}
			resp.Body.Close()
		}
		for _, v := range schedulerMonitor { //todo：面对实时修改的map使用range可能会出问题
			tempCountCpu += v.cpuload
			tempCountMem += v.memload
		}
		if tempCountCpu*10 > len(schedulerMonitor)*(loadFactorCpu*10) {
			// 发起HTTP GET请求
			resp, err := http.Get(gatewayAddr + ":" + constant.GatewayHttpPort + "/create?" + "funcName=scheduler&requireCpu=4&requireMem=4")
			if err != nil {
				fmt.Errorf("扩容失败:", err)
			}
			resp.Body.Close()
		}
		for _, v := range nodeMonitor { //todo：面对实时修改的map使用range可能会出问题
			tempCountCpu += v.cpuload
			tempCountMem += v.memload
		}
		if tempCountCpu*10 > len(nodeMonitor)*(loadFactorCpu*10) {
			// 发起HTTP GET请求
			resp, err := http.Get(gatewayAddr + ":" + constant.GatewayHttpPort + "/create?" + "funcName=node&requireCpu=4&requireMem=4")
			if err != nil {
				fmt.Errorf("扩容失败:", err)
			}
			resp.Body.Close()
		}
	}
}
