package main

import (
	"flag"
	pb "github.com/ServerlessOS/galaxy/proto"
	"google.golang.org/grpc"
	"net"
	"virtualNode_rpc/server"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "p", "16446", "启动端口号")
	flag.Parse()
}

func main() {
	//logFile, err := os.OpenFile("./logs/QPS200.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatal("Error opening log file:", err)
	//}
	//defer logFile.Close()
	//log.SetOutput(logFile)

	go server.InstanceInfoInform()
	s := grpc.NewServer()
	pb.RegisterNodeServer(s, &server.NodeServer{})

	lis, _ := net.Listen("tcp", ":"+port)
	s.Serve(lis)
}
