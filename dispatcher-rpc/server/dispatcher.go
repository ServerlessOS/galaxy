package server

import (
	"context"
	"dispatcher_rpc/internal"
	"fmt"
	pb "github.com/ServerlessOS/galaxy/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var (
	funcView      = internal.NewFuncView() // 支持查找变更 func
	SchedulerView = internal.NewSchedulerView()
	localIp       = getLocalIPv4().String()
	connQueue     = NewConnQueue()
	firstInit     = true
)

type DispatcherServer struct{}

func (d DispatcherServer) Statis(ctx context.Context, request *pb.UserRequest) (*pb.UserRequestReply, error) {
	return &pb.UserRequestReply{
		RequestId:   1,
		FuncName:    "request.FuncName",
		Destination: "result",
	}, nil
}

func (d DispatcherServer) UpdateInstanceView(ctx context.Context, update *pb.InstanceUpdate) (*pb.InstanceUpdateReply, error) {
	list := update.List
	for _, v := range list {
		t := time.Now().UnixNano() / 1e6
		item := &internal.FuncInfo{FuncName: v.FuncName,
			Address:   v.Address,
			Timestamp: t,
			State:     true}
		if update.Action == "ADD" {
			funcView.Add(item)
			log.Printf("Add new instance -%d-%s-%s at -%d\n", v.RequestId, item.FuncName, item.Address, time.Now().UnixNano()/1e6)
			//fmt.Printf("Add new instance -%d-%s-%s at -%d\n", v.RequestId, item.FuncName, item.Address, time.Now().UnixNano()/1e6)
		} else if update.Action == "DELETE" {
			funcView.Delete(item)
		}
	}
	return &pb.InstanceUpdateReply{
		State: 0,
	}, nil
}
func randonInitConnection() {
	for i := 0; i < 20; i++ {
		saddr := SchedulerView.GetSchedulerAddr()
		conn, _ := grpc.Dial(fmt.Sprintf("%s:16445", saddr), grpc.WithInsecure())
		connQueue.Enqueue(conn)
	}
}
func addConn(n int) {
	for i := 0; i < n; i++ {
		saddr := SchedulerView.GetSchedulerAddr()
		conn, _ := grpc.Dial(fmt.Sprintf("%s:16445", saddr), grpc.WithInsecure())
		connQueue.Enqueue(conn)
	}
}
func (d DispatcherServer) Dispatch(ctx context.Context, userRequests *pb.UserRequestList) (*pb.UserRequestReply, error) {
	t := time.Now().UnixNano() / 1e6
	per_scheduler := len(userRequests.List) / SchedulerView.GetLen()
	per_scheduler = per_scheduler / 2
	schedulerRequests := make([]*pb.ScheduleRequest, 0)
	for _, request := range userRequests.List {
		result := funcView.Dispatch(request.FuncName)
		if result == "" {
			log.Printf("Need to scale up -%d-%s at -%d\n", request.RequestId, request.FuncName, t)
			//lock.Lock()
			//requestIdList = append(requestIdList, request.RequestId)
			//lock.Unlock()
			schedulerRequests = append(schedulerRequests, &pb.ScheduleRequest{
				RequestId:      request.RequestId,
				FuncName:       request.FuncName,
				RequireCpu:     request.RequireCpu,
				RequireMem:     request.RequireMem,
				DispatcherAddr: localIp,
			})
			//fmt.Printf("len:%d\n", len(schedulerRequests))
			if len(schedulerRequests) >= per_scheduler {
				c := connQueue.Dequeue()
				client := pb.NewSchedulerClient(c)
				//fmt.Printf("Route %d requests to %s\n",len(schedulerRequests),c.GetMethodConfig().)
				_, _ = client.Schedule(context.Background(), &pb.ScheduleRequestList{List: schedulerRequests})
				go addConn(2)
				schedulerRequests = make([]*pb.ScheduleRequest, 0)
			}
		}
	}
	if len(schedulerRequests) > 0 {
		c := connQueue.Dequeue()
		client := pb.NewSchedulerClient(c)
		_, _ = client.Schedule(context.Background(), &pb.ScheduleRequestList{List: schedulerRequests})
		go addConn(2)
		//lock.Lock()
		//fmt.Printf("cnt: %d\n", uniqueKeyCount(requestIdList))
		//lock.Unlock()
	}

	return &pb.UserRequestReply{
		RequestId:   0,
		FuncName:    "0",
		Destination: "0",
	}, nil
}
func (d DispatcherServer) UpdateSchedulerView(ctx context.Context, update *pb.SchedulerViewUpdate) (*pb.SchedulerViewUpdateReply, error) {
	//fmt.Printf("Add new scheduelr \n")
	list := update.List
	for _, v := range list {
		item := &internal.SchedulerInfo{NodeName: v.NodeName, Address: v.Address}
		if update.Action == "ADD" {
			fmt.Printf("Add new scheduelr %s:%s\n", item.NodeName, item.Address)
			SchedulerView.Add(item)

		} else if update.Action == "DELETE" {
			SchedulerView.Delete(item)
		}
	}
	if firstInit {
		go addConn(20)
		firstInit = false
	}
	return &pb.SchedulerViewUpdateReply{State: 0}, nil
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
func uniqueKeyCount(arr []int64) int {
	uniqueKeys := make(map[int64]struct{})

	for _, item := range arr {
		uniqueKeys[item] = struct{}{}
	}

	return len(uniqueKeys)
}
