package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/client"
	gateway_rpc "github.com/ServerlessOS/galaxy/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	localHttpAddr        string
	localRpcAddr         string
	coordinatorAddr      string
	gatewayName          string
	gatewayList          = make(map[string]string)
	gatewayList_mutex    sync.Mutex
	dispatcherList       = make(map[string]string)
	dispatcherList_mutex sync.Mutex
	funcManagerList      = make(map[string]string)
	funcManager_mutex    sync.Mutex
	localIp              = getLocalIPv4().String()
)
var Cmd = &cobra.Command{
	Use:   "gateway",
	Short: `初始化gateway程序`,
	//本函数用于执行命令并返回错误
	RunE: func(cmd *cobra.Command, args []string) error {
		var errChanHttp, errChanRpc chan error
		go httpServer(errChanHttp)
		go rpcServer(errChanRpc)
		initGateway()
		select {
		case errHttp := <-errChanHttp:
			return errHttp
		case errRpc := <-errChanRpc:
			return errRpc
		}
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

// NewGatewayCommand 生成根命令，启动服务
func init() {
	Cmd.Flags().StringVarP(&localHttpAddr, "localHttpAddr", "p", ":16447", "The addr used for binding to the HTTP server. ")
	Cmd.Flags().StringVarP(&localRpcAddr, "localRpcAddr", "r", ":16448", "The addr used for binding to the RPC server. ")
	Cmd.Flags().StringVarP(&coordinatorAddr, "coordinatorAddr", "c", "", "The addr used for connect to the coordinator. ")
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
}
func initGateway() {
	//连接顶层控制器，注册
	gatewayName = strconv.Itoa(int(rand.Uint32()))
	err := client.DialCoordinatorClient("def", coordinatorAddr)

	if err != nil {
		log.Fatalln("client err", err)
	}
	log.Println("coordinator connect success")
	//既可以让gateway自己向顶层控制器注册，也可以经由其它gateway向顶层控制器注册，为了保证gateway0和gateway1做法一致，所以采用自行注册的方案
	if isIPAddress(localRpcAddr) {
		tcpAddr, err := net.ResolveTCPAddr("tcp", localRpcAddr)
		if err != nil {
			log.Fatalln(err)
		}
		if tcpAddr.IP.String() == "<nil>" { //todo:!这里的判断有问题
			client.GetCoordinatorClient().Register(context.Background(), &gateway_rpc.RegisterReq{
				Type:    0,
				Name:    gatewayName,
				Address: localIp,
			})
			log.Println("register gateway,ip:", localIp)
		} else {
			client.GetCoordinatorClient().Register(context.Background(), &gateway_rpc.RegisterReq{
				Type:    0,
				Name:    gatewayName,
				Address: tcpAddr.IP.String(),
			})
			log.Println("register gateway,ip:", tcpAddr.IP.String())
		}
	} else {
		client.GetCoordinatorClient().Register(context.Background(), &gateway_rpc.RegisterReq{
			Type:    0,
			Name:    gatewayName,
			Address: localRpcAddr,
		})
		log.Println("register gateway,ip:", localRpcAddr)
		if err != nil {
			log.Fatalln(err)
		}
	}

}

func httpServer(errChannel chan<- error) {
	//gateway与上游DNS服务器对接与扩容
	http.HandleFunc("/getGatewayList", getGatewayList)

	http.HandleFunc("/create", create)
	log.Println("http address:", localHttpAddr)
	err := http.ListenAndServe(localHttpAddr, nil)
	errChannel <- err
}
func rpcServer(errChannel chan<- error) {
	lis, err := net.Listen("tcp", localRpcAddr)
	if err != nil {
		errChannel <- err
	}

	// 实例化grpc服务端
	s := grpc.NewServer()

	// 在gRPC服务器注册服务
	gateway_rpc.RegisterGatewayServer(s, &rpcServerProcess{})

	// 启动grpc服务
	err = s.Serve(lis)
	errChannel <- err
}

type rpcServerProcess struct{}

func (r rpcServerProcess) UpdateGatewayList(ctx context.Context, req *gateway_rpc.UpdateListReq) (*gateway_rpc.UpdateListResp, error) {
	//    APPEND = 0;
	//    REDUCE = 1;
	//    OVERRIDE = 2;
	switch req.Type {
	case 0:
		for k, v := range req.List {
			gatewayList[k] = v
		}
	case 1:
		for k, _ := range req.List {
			delete(gatewayList, k)
		}
	case 2:
		gatewayList = req.List
	default:
		return &gateway_rpc.UpdateListResp{
			StatusCode:  1,
			Description: "undefined operation type",
		}, fmt.Errorf("undefined operation type")
	}
	return &gateway_rpc.UpdateListResp{
		StatusCode:  0,
		Description: "OK",
	}, nil
}

func (r rpcServerProcess) UpdateDispatcherList(ctx context.Context, req *gateway_rpc.UpdateListReq) (*gateway_rpc.UpdateListResp, error) {
	//    APPEND = 0;
	//    REDUCE = 1;
	//    OVERRIDE = 2;
	switch req.Type {
	case 0:
		for k, v := range req.List {
			dispatcherList[k] = v
		}
	case 1:
		for k, _ := range req.List {
			delete(dispatcherList, k)
		}
	case 2:
		dispatcherList = req.List
	default:
		return &gateway_rpc.UpdateListResp{
			StatusCode:  1,
			Description: "undefined operation type",
		}, fmt.Errorf("undefined operation type")
	}
	return &gateway_rpc.UpdateListResp{
		StatusCode:  0,
		Description: "OK",
	}, nil
}

func (r rpcServerProcess) UpdateFuncManagerList(ctx context.Context, req *gateway_rpc.UpdateListReq) (*gateway_rpc.UpdateListResp, error) {
	//    APPEND = 0;
	//    REDUCE = 1;
	//    OVERRIDE = 2;
	switch req.Type {
	case 0:
		for k, v := range req.List {
			funcManagerList[k] = v
		}
	case 1:
		for k, _ := range req.List {
			delete(funcManagerList, k)
		}
	case 2:
		funcManagerList = req.List
	default:
		return &gateway_rpc.UpdateListResp{
			StatusCode:  1,
			Description: "undefined operation type",
		}, fmt.Errorf("undefined operation type")
	}
	return &gateway_rpc.UpdateListResp{
		StatusCode:  0,
		Description: "OK",
	}, nil
}

// 将注册请求转发给顶层控制器
func (r rpcServerProcess) Register(ctx context.Context, req *gateway_rpc.RegisterReq) (*gateway_rpc.RegisterResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if req.Address == "" {
		pr, ok := peer.FromContext(ctx)
		if !ok {
			return nil, fmt.Errorf("无法获取客户端信息")
		}
		// 获取客户端IP地址
		addr := pr.Addr
		tcpAddr, ok := addr.(*net.TCPAddr)
		if !ok {
			return nil, fmt.Errorf("无法获取客户端IP地址")
		}
		req.Address = tcpAddr.IP.String()
	}
	resp, err := client.GetCoordinatorClient().Register(ctx, req)
	return resp, err
}

func create(w http.ResponseWriter, req *http.Request) {
	// 解析 URL 中的查询参数
	queryParams := req.URL.Query()
	// 获取特定参数的值
	funcName := queryParams.Get("funcName")
	requireCpuString := queryParams.Get("requireCpu")
	requireMemString := queryParams.Get("requireMem")
	requireCpu, _ := strconv.Atoi(requireCpuString)
	requireMem, _ := strconv.Atoi(requireMemString)
	//调用dispatcher
	_, err := client.GetDispatcherClient().Dispatch(context.Background(), &gateway_rpc.UserRequestList{List: []*gateway_rpc.UserRequest{
		{RequestId: rand.Int63(), FuncName: funcName, RequireCpu: int64(requireCpu), RequireMem: int64(requireMem)},
	}})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("dispatcher success.")
	// 发送成功的 HTTP 响应
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("create success."))
}
func getGatewayList(w http.ResponseWriter, req *http.Request) {
	listString, err := json.Marshal(gatewayList)
	if err != nil {
		log.Errorln(err)
	}
	w.Write(listString)
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
