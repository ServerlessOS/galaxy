package main

//
//import (
//	"context"
//	pb "github.com/ServerlessOS/galaxy/proto"
//	"flag"
//	"fmt"
//	"google.golang.org/grpc"
//	"log"
//	"time"
//)
//
//var port string
//
//func init() {
//	flag.StringVar(&port, "p", "16444", "启动端口号")
//	flag.Parse()
//}
//
//func main() {
//	//init Scheduler
//	conn, _ := grpc.Dial(":16445", grpc.WithInsecure())
//	defer conn.Close()
//	client := pb.NewSchedulerClient(conn)
//	_ = updateResource(client)
//	time.Sleep(2 * time.Second)
//	//init dispatcher
//	connDispatcher, _ := grpc.Dial(":"+port, grpc.WithInsecure())
//	defer connDispatcher.Close()
//	clientDispatcher := pb.NewDispatcherClient(connDispatcher)
//	//updateInstance(clientDispatcher)
//	time.Sleep(1 * time.Second)
//	updateScheduler(clientDispatcher)
//	time.Sleep(1 * time.Second)
//	getForward(clientDispatcher)
//	//time.Sleep(1 * time.Second)
//	//_, _ = clientDispatcher.Statis(context.Background(), &pb.UserRequest{
//	//	RequestId: int64(1*10 + 2),
//	//	FuncName:  fmt.Sprintf("Func-%d", 2),
//	//})
//}
//func updateInstance(client pb.DispatcherClient) {
//	list := make([]*pb.InstanceInfo, 0)
//	for i := 0; i < 10; i++ {
//		list = append(list, &pb.InstanceInfo{
//			FuncName: fmt.Sprintf("Func-%d", i),
//			Address:  fmt.Sprintf("192.168.1.%d", i),
//		})
//	}
//	resp, _ := client.UpdateInstanceView(context.Background(), &pb.InstanceUpdate{List: list, Action: "ADD"})
//	log.Printf("client.UpadateNodeResource resp: %d", resp.State)
//}
//func updateScheduler(client pb.DispatcherClient) {
//	list := make([]*pb.SchedulerInfo, 0)
//	list = append(list, &pb.SchedulerInfo{
//		NodeName: fmt.Sprintf("Scheduler-%d", 1),
//		Address:  fmt.Sprintf("127.0.0.1"),
//	})
//	resp, _ := client.UpdateSchedulerView(context.Background(), &pb.SchedulerViewUpdate{List: list, Action: "ADD"})
//	log.Printf("client.UpadateNodeResource resp: %d", resp.State)
//}
//func getForward(client pb.DispatcherClient) {
//	//message UserRequest{
//	//	int64 RequestId = 1;
//	//	string FuncName = 2;
//	//}
//	for i := 0; i < 2; i++ {
//
//		ur := &pb.UserRequest{
//			RequestId: int64(i),
//			FuncName:  fmt.Sprintf("Func-%d", i),
//		}
//		resp, _ := client.Dispatch(context.Background(), ur)
//		if resp.Destination != "" {
//			log.Printf("client.scale resp: %s", resp.Destination)
//		}
//
//	}
//}
//func updateResource(client pb.SchedulerClient) error {
//	list := make([]*pb.NodeResource, 0)
//	for i := 0; i < 10; i++ {
//		list = append(list, &pb.NodeResource{
//			NodeName: fmt.Sprintf("node-%d", i),
//			HaveCpu:  int64(i * 100),
//			HaveMem:  int64(i * 100),
//			Address:  "127.0.0.1",
//		})
//	}
//	resp, _ := client.UpadateNodeResource(context.Background(), &pb.NodeResourceUpdate{List: list, Action: "ADD"})
//	log.Printf("client.UpadateNodeResource resp: %d", resp.State)
//	return nil
//}
