package register

import (
	"context"
	assignor "coordinator_rpc/RendezousHashing"
	"coordinator_rpc/cmd"
	"coordinator_rpc/constant"
	pb "coordinator_rpc/proto"
	"coordinator_rpc/server"
	"log"
	"time"
)

var (
	SchedulerList = make([]*pb.SchedulerInfo, 0)
)

type Scheduler struct {
	RootModule
}

func (s *Scheduler) Register() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	sch := assignor.NewScheduler(s.Pod.Name, s.getPodIP())
	server.Rh.Schedulers[sch.Name] = sch
	SchedulerList = append(SchedulerList, sch.ToProto())
	cmd.DialSchedulerClient(sch.Name, sch.Addr+":"+constant.SchedulerPort)
	for nodeName, nodeResource := range server.Rh.Nodes {
		hash := server.Rh.Hash(sch.Name + nodeName)
		if hash > nodeResource.Hash {
			oldScheduler := server.Rh.Schedulers[server.Rh.SNView[nodeName]]
			// remove the node from the schueduler
			req := &pb.NodeResourceUpdate{
				List:       []*pb.NodeResource{nodeResource.ToProto()},
				Action:     "DELETE",
				SourceAddr: oldScheduler.Addr,
				TargetAddr: sch.Addr,
			}
			client := cmd.GetSchedulerClient(oldScheduler.Name)
			resp, err := client.UpadateNodeResource(ctx, req)
			if resp.State != 0 || err != nil {
				panic(err)
			}
			nodeResource.Hash = hash
			server.Rh.SNView[nodeName] = sch.Name
		}
	}
	//通知上游的dispatcher
	dClientList := cmd.GetDispatcherClientList()
	for _, client := range dClientList {
		resp, _ := client.UpdateSchedulerView(ctx, &pb.SchedulerViewUpdate{
			List: []*pb.SchedulerInfo{{NodeName: sch.Name, Address: sch.Addr}}, Action: "ADD",
		})
		log.Printf("client.UpdateSchedulerView resp: %d", resp.State)
	}

	//通告新scheduler出现
	sClientList := cmd.GetSchedulerClientList()
	for schedulerName, client := range sClientList {
		if schedulerName != sch.Name {
			resp, _ := client.PeerSchedulerUpdate(ctx, &pb.PeerSchedulersUpdate{
				List:   []*pb.PeerSchedulerInfo{{NodeName: sch.Name, Address: sch.Addr}},
				Action: "ADD",
			})
			log.Printf("client.UpadateNodeResource resp: %d", resp.State)
		}
	}

	//对新scheduler同步旧scheduler的信息
	schClient := cmd.GetSchedulerClient(sch.Name)
	peerSchedulerlist := []*pb.PeerSchedulerInfo{}
	for _, scheduler := range SchedulerList {
		if scheduler.NodeName != sch.Name {
			peerSchedulerlist = append(peerSchedulerlist, &pb.PeerSchedulerInfo{
				NodeName: scheduler.NodeName,
				Address:  scheduler.Address,
			})
		}
	}
	resp, err := schClient.PeerSchedulerUpdate(ctx, &pb.PeerSchedulersUpdate{
		List:   peerSchedulerlist,
		Action: "ADD",
	})
	if err != nil {
		log.Fatalln("PeerSchedulerUpdate err,", err)
	}
	log.Printf("register scheduler, name:%v,state:%s", sch.Name, resp.State)
}