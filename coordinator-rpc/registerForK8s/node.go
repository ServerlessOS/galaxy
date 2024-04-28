package registerForK8s

import (
	"context"
	assignor "coordinator_rpc/RendezousHashing"
	"coordinator_rpc/client"
	"coordinator_rpc/register"
	"github.com/ServerlessOS/galaxy/constant"
	pb "github.com/ServerlessOS/galaxy/proto"
	"log"
	"time"
)

var (
	nodeCache = make([]*Node, 0)
)

type Node struct {
	RootModule
}

func (n *Node) Register() {
	//如果此时没有scheduler，就先缓存后续再处理
	if len(register.Rh.Schedulers) == 0 {
		nodeCache = append(nodeCache, n)
		return
	}
	//准备待发送的数据
	name, address := n.Pod.Name, n.getPodIP()
	node := &assignor.NodeResource{
		NodeName: name,
		HaveCpu:  constant.NodeCpu,
		HaveMem:  constant.NodeMem,
		Address:  address,
		Port:     constant.NodePort,
		Hash:     0,
	}
	list := []*pb.NodeResource{node.ToProto()}

	//挑选数据接收者
	var selectedScheduler *assignor.Scheduler
	maxHash := uint32(0)
	for _, scheduler := range register.Rh.Schedulers {
		hash := register.Rh.Hash(scheduler.Name + node.NodeName)
		if hash > maxHash {
			maxHash = hash
			selectedScheduler = scheduler
		}
	}
	node.Hash = assignor.FnvHash(node.NodeName)
	register.Rh.Nodes[node.NodeName] = node
	register.Rh.SNView[node.NodeName] = selectedScheduler.Name

	sClient := client.GetSchedulerClient(selectedScheduler.Name)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := sClient.UpadateNodeResource(ctx, &pb.NodeResourceUpdate{List: list, Action: "ADD"})
	if err != nil {
		log.Fatalln("UpadateNodeResource err,name:", name, ",err:", err)
	}
	log.Printf("registerForK8s node, name:%v,state:%s", name, resp.State)
	if len(nodeCache) != 0 {
		//曾经缓存的注册请求
		node := nodeCache[0]
		nodeCache = nodeCache[1:]
		node.Register()
	}
}
