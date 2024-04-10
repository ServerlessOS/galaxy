package main

import (
	_ "coordinator_rpc/informer"
	"coordinator_rpc/server"
	"flag"
	pb "github.com/ServerlessOS/galaxy/proto"
	"net"

	"google.golang.org/grpc"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "p", "16000", "启动端口号")
	flag.Parse()
}
func main() {
	s := grpc.NewServer()
	pb.RegisterCoordinatorServer(s, &server.CoordiantorServer{})
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
