package main

import (
	"flag"
	"google.golang.org/grpc"
	"net"
	pb "scheduler_rpc/proto"
	"scheduler_rpc/server"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "p", "16445", "启动端口号")
	flag.Parse()

}

func main() {
	//logFile, err := os.OpenFile("./logs/QPS200.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatal("Error opening log file:", err)
	//}
	//defer logFile.Close()
	//log.SetOutput(logFile)

	//go server.ResourceUpdate()
	go server.Schedule()
	// 启动rpc server
	s := grpc.NewServer()
	pb.RegisterSchedulerServer(s, &server.SchedulerServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	s.Serve(lis)
}
