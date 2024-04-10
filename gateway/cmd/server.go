package cmd

import (
	"context"
	gateway_rpc "github.com/ServerlessOS/galaxy/proto"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	localHttpAddr   string
	localRpcAddr    string
	coordinatorAddr string
	gatewayName     string
)
var Cmd = &cobra.Command{
	Use:   "gateway",
	Short: `初始化gateway程序`,
	//本函数用于执行命令并返回错误
	RunE: func(cmd *cobra.Command, args []string) error {
		gatewayName = strconv.Itoa(int(rand.Uint32()))
		err := clientInit()
		if err != nil {
			log.Fatalln("client err", err)
		}
		var errChanHttp, errChanRpc chan error
		go httpServer(errChanHttp)
		go rpcServer(errChanRpc)
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

func httpServer(errChannel chan<- error) {
	//gateway与上游DNS服务器对接与扩容
	http.HandleFunc("/getGatewayList", getGatewayList)
	http.HandleFunc("/extensionGateway", extensionGateway)

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

func (r rpcServerProcess) UpdateGatewayList(ctx context.Context, req *gateway_rpc.UpdateGatewayListReq) (*gateway_rpc.UpdateGatewayListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (r rpcServerProcess) Register(ctx context.Context, req *gateway_rpc.RegisterReq) (*gateway_rpc.RegisterResp, error) {
	//TODO implement me
	panic("implement me")
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
	_, err := GetDispatcherClient().Dispatch(context.Background(), &gateway_rpc.UserRequestList{List: []*gateway_rpc.UserRequest{
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
