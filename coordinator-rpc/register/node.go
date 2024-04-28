package register

import (
	"context"
	assignor "coordinator_rpc/RendezousHashing"
	"coordinator_rpc/client"
	"fmt"

	"github.com/ServerlessOS/galaxy/constant"
	pb "github.com/ServerlessOS/galaxy/proto"
	"log"
	"time"
)

type Node struct {
}

func (n *Node) Register(req *pb.RegisterReq) error {
	if len(Rh.Schedulers) == 0 {
		return fmt.Errorf("not have scheduler can be choose")
	}
	//准备待发送的数据
	name, address := req.Name, req.Address
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
	for _, scheduler := range Rh.Schedulers {
		hash := Rh.Hash(scheduler.Name + node.NodeName)
		if hash > maxHash {
			maxHash = hash
			selectedScheduler = scheduler
		}
	}
	node.Hash = assignor.FnvHash(node.NodeName)
	Rh.Nodes[node.NodeName] = node
	Rh.SNView[node.NodeName] = selectedScheduler.Name

	sClient := client.GetSchedulerClient(selectedScheduler.Name)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := sClient.UpadateNodeResource(ctx, &pb.NodeResourceUpdate{List: list, Action: "ADD"})
	if err != nil {
		log.Fatalln("UpadateNodeResource err,name:", name, ",err:", err)
	}
	log.Printf("registerForK8s node, name:%v,state:%s", name, resp.State)
	return nil
}
