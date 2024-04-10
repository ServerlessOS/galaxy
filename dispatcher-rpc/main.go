package main

import (
	"dispatcher_rpc/server"
	"flag"
	pb "github.com/ServerlessOS/galaxy/proto"
	"google.golang.org/grpc"
	"net"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "p", "16444", "启动端口号")
	flag.Parse()
}

func main() {
	s := grpc.NewServer()
	pb.RegisterDispatcherServer(s, &server.DispatcherServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	s.Serve(lis)
}
