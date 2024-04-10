package main

//
//import (
//	"context"
//	"flag"
//	"fmt"
//	"google.golang.org/grpc"
//	"log"
//	pb "github.com/ServerlessOS/galaxy/proto"
//	"time"
//)
//
//var port string
//
//func init() {
//	flag.StringVar(&port, "p", "16445", "启动端口号")
//	flag.Parse()
//}
//
//func main() {
//	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
//	defer conn.Close()
//	client := pb.NewSchedulerClient(conn)
//
//	_ = updateResource(client)
//	time.Sleep(2 * time.Second)
//	_ = scaleRequest(client)
//
//}
//func updateResource(client pb.SchedulerClient) error {
//	list := make([]*pb.NodeResource, 0)
//	for i := 0; i < 10; i++ {
//		list = append(list, &pb.NodeResource{
//			NodeName: fmt.Sprintf("node-%d", i),
//			HaveCpu:  int64(i * 100),
//			HaveMem:  int64(i * 100),
//		})
//	}
//	resp, _ := client.UpadateNodeResource(context.Background(), &pb.NodeResourceUpdate{List: list, Action: "ADD"})
//	log.Printf("client.UpadateNodeResource resp: %d", resp.State)
//	return nil
//}
//func scaleRequest(client pb.SchedulerClient) error {
//	for i := 0; i < 10; i++ {
//		request := &pb.ScheduleRequest{
//			RequestId:  int64(i),
//			FuncName:   fmt.Sprintf("func-%d", i),
//			RequireCpu: int64(i * 20),
//			RequireMem: int64(i * 20),
//		}
//		resp, _ := client.Schedule(context.Background(), request)
//		log.Printf("client.scaleRequest funcId: %d", resp.RequestId)
//	}
//	return nil
//}
