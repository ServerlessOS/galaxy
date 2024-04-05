package cmd

import (
	"context"
	"fmt"
	"func-manager/proto"
	"net"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func Run(cmd *cobra.Command) (code int) {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	return 0
}

// todo:子命令、args解析、文档完善
func NewGatewayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "gateway",
		Long: ``,
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	return cmd
}

func rpcServer(addr string, errChannel chan<- error) {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		errChannel <- err
	}
	// 实例化grpc服务端
	s := grpc.NewServer()

	// 在gRPC服务器注册服务
	proto.RegisterFuncMagagerServer(s, &rpcServerProcess{})

	// 启动grpc服务
	err = s.Serve(lis)
	errChannel <- err
}

type rpcServerProcess struct{}

// 实现Create接口
func (s *rpcServerProcess) Create(ctx context.Context, req *proto.CreateReq) (*proto.GeneralResp, error) {
	return &proto.GeneralResp{}, nil
}

// 实现Register接口
func (s *rpcServerProcess) Register(ctx context.Context, in *proto.RegisterReq) (*proto.GeneralResp, error) {
	// 创建一个消息，设置Message字段，然后直接返回。
	return &proto.GeneralResp{}, nil
}

// 实现Delete接口
func (s *rpcServerProcess) Delete(ctx context.Context, req *proto.DeleteReq) (*proto.GeneralResp, error) {
	return &proto.GeneralResp{}, nil
}

// 实现Get接口
func (s *rpcServerProcess) Get(ctx context.Context, req *proto.GetReq) (*proto.GetResp, error) {
	return &proto.GeneralResp{}, nil
}

// 实现Describe接口
func (s *rpcServerProcess) Describe(ctx context.Context, req *proto.DescribeReq) (*proto.DescribeResp, error) {
	return &proto.GeneralResp{}, nil
}

// 实现Logs接口
func (s *rpcServerProcess) Logs(ctx context.Context, req *proto.LogsReq) (*proto.LogsResp, error) {
	return &proto.GeneralResp{}, nil
}

// 实现Version接口
func (s *rpcServerProcess) Version(ctx context.Context, req *proto.VersionReq) (*proto.VersionResp, error) {
	return &proto.GeneralResp{}, nil
}

// 实现Label接口
func (s *rpcServerProcess) Label(ctx context.Context, req *proto.LabelReq) (*proto.GeneralResp, error) {
	return &proto.GeneralResp{}, nil
}

// 实现Annotation接口
func (s *rpcServerProcess) Annotation(ctx context.Context, req *proto.AnnotationReq) (*proto.GeneralResp, error) {
	return &proto.GeneralResp{}, nil
}
