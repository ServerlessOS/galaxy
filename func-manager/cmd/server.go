package cmd

import (
	"context"
	"fmt"
	"github.com/ServerlessOS/galaxy/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"math/rand/v2"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var (
	localRpcAddr    string
	coordinatorAddr string
	funcManagerName string
	funcYaml        map[string]*function //通过函数名索引函数，todo：考虑持久化
	ExitCh          = make(chan int)
)

type function struct {
	Name       string
	Label      string
	Annotation string
	Document   string
}

var Cmd = &cobra.Command{
	Use:   "gateway",
	Short: `初始化gateway程序`,
	//本函数用于执行命令并返回错误
	RunE: func(cmd *cobra.Command, args []string) error {
		funcManagerName = strconv.Itoa(int(rand.Uint32()))
		var errChanRpc chan error
		rpcServer(errChanRpc)
		err := <-errChanRpc
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
			return err
		}
		return nil
	},
}

func init() {
	Cmd.Flags().StringVarP(&localRpcAddr, "localRpcAddr", "r", ":16449", "The addr used for binding to the RPC server. ")
	Cmd.Flags().StringVarP(&coordinatorAddr, "coordinatorAddr", "c", "", "The addr used for connect to the coordinator. ")
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
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

func rpcServer(errChannel chan<- error) {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		errChannel <- err
	}
	// 实例化grpc服务端
	s := grpc.NewServer()

	// 在gRPC服务器注册服务
	proto.RegisterFuncManagerServer(s, &rpcServerProcess{})

	// 启动grpc服务
	err = s.Serve(lis)
	errChannel <- err
}
func GracefulExit() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Kill)

	sig := <-signalChan
	log.Printf("catch signal, %+v", sig)

	file, err := yaml.Marshal(funcYaml)
	if err != nil {
		log.Fatal("yaml conversion error")
	}
	if err := os.WriteFile("data.txt", file, 0644); err != nil {
		log.Fatal("file save err:", err)
	}
	log.Printf("data persistence successful")
	close(ExitCh)
}

type rpcServerProcess struct{}

func (s *rpcServerProcess) Create(ctx context.Context, req *proto.CreateReq) (*proto.CreateResp, error) {
	//todo：成熟时应该添加某种程度的校验和鉴权
	newFunc := &function{
		Name:       req.Request.Name,
		Label:      req.Request.Labels,
		Annotation: req.Annotations,
		Document:   req.Document,
	}
	funcYaml[newFunc.Name] = newFunc
	return &proto.CreateResp{
		RequestId:        req.Request.RequestId,
		StatusCode:       0,
		Description:      "OK",
		ErrorInformation: "",
	}, nil
}

func (s *rpcServerProcess) Get(ctx context.Context, req *proto.GetReq) (*proto.GetResp, error) {
	targetFunc, ok := funcYaml[req.Request.Name]
	if !ok {
		return nil, fmt.Errorf("functione not found")
	}
	return &proto.GetResp{
		RequestId:  req.Request.RequestId,
		StatusCode: 0,
		Document:   targetFunc.Document,
	}, nil

}
func (s *rpcServerProcess) Delete(ctx context.Context, req *proto.DeleteReq) (*proto.DeleteResp, error) {
	_, ok := funcYaml[req.Request.Name]
	if !ok {
		return nil, fmt.Errorf("functione not found")
	}
	delete(funcYaml, req.Request.Name)
	return &proto.DeleteResp{
		RequestId:        req.Request.RequestId,
		StatusCode:       0,
		Description:      "OK",
		ErrorInformation: "",
	}, nil
}

func (s *rpcServerProcess) List(ctx context.Context, req *proto.ListReq) (*proto.ListResp, error) {
	var funcName []string
	funcName = maps.Keys(funcYaml)
	yamlData, err := yaml.Marshal(funcName)
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, fmt.Errorf("yaml conversion error")
	}
	return &proto.ListResp{
		RequestId:  req.RequestId,
		StatusCode: 0,
		List:       string(yamlData),
	}, nil
}
