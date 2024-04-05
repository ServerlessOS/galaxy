package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
	"virtualNode_rpc/internal"
	pb "virtualNode_rpc/proto"
)

var (
	DeployQueue = internal.NewFIFO()
	ConnCache   *LRUCache
)

func init() {
	ConnCache = NewLRUCache(20)
}

type NodeServer struct{}

func (n NodeServer) Deploy(ctx context.Context, deploy *pb.InstanceDeploy) (*pb.InstanceDeployReply, error) {
	list := deploy.List
	for _, instanceInfo := range list {
		newInstance := &internal.InstanceInfo{
			RequestId:         instanceInfo.RequestId,
			FuncName:          instanceInfo.FuncName,
			Address:           randomIP().String(),
			DispatcherAddress: instanceInfo.DispatcherAddr,
		}
		DeployQueue.Enqueue(newInstance)
		log.Printf("Deploy instace -%d-%s at -%d\n", instanceInfo.RequestId, instanceInfo.FuncName, time.Now().UnixNano()/1e6)
		//fmt.Printf("Deploy instace -%d-%s at -%d\n", instanceInfo.RequestId, instanceInfo.FuncName, time.Now().UnixNano()/1e6)
	}

	return &pb.InstanceDeployReply{State: 0}, nil
}

func InstanceInfoInform() {
	for {
		instanceInfo := DeployQueue.Dequeue()
		dispatcherAddr := instanceInfo.DispatcherAddress
		// dispatcher 端client调用
		ctx := context.Background()
		conn := ConnCache.Get(dispatcherAddr)
		if conn != nil {
			fmt.Printf("Cache hitted!\n")
		} else {
			conn, _ = grpc.Dial(fmt.Sprintf("%s:16444", dispatcherAddr), grpc.WithInsecure())
			ConnCache.Put(dispatcherAddr, conn)
			fmt.Printf("Cache miss!\n")
		}
		client := pb.NewDispatcherClient(conn)
		_, err := client.UpdateInstanceView(ctx, &pb.InstanceUpdate{
			List: []*pb.InstanceInfo{
				&pb.InstanceInfo{
					RequestId: instanceInfo.RequestId,
					FuncName:  instanceInfo.FuncName,
					Address:   instanceInfo.Address,
				},
			},
			Action: "ADD",
		})
		if err != nil {
			return
		}
		log.Printf("Node update instance -%d- at -%d\n", instanceInfo.RequestId, time.Now().UnixNano()/1e6)
	}
}
func randomIP() net.IP {
	// 生成随机的IPv4地址
	ip := make(net.IP, 4)
	rand.Seed(time.Now().UnixNano())
	for i := range ip {
		ip[i] = byte(rand.Intn(256))
	}
	return ip
}
