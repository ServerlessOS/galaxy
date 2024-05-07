package register

import (
	"context"
	assignor "coordinator_rpc/RendezousHashing"
	"coordinator_rpc/client"
	pb "github.com/ServerlessOS/galaxy/proto"
	log "github.com/sirupsen/logrus"
	"time"
)

type ClusterManager struct {
}

func (d *ClusterManager) Register(req *pb.RegisterReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	name, address := req.Name, req.Address
	log.Println("new clusterManager,name:", name, "\naddress:", address)
	clus := &assignor.ClusterManager{
		Name: name,
		Addr: address,
	}
	Rh.ClusterManager[clus.Name] = clus

	//给gateway同步Cluster Manager
	for s, _ := range Rh.Gateways {
		resp1, err := client.GetGatewayClient(s).UpdateClusterManagerList(ctx, &pb.UpdateListReq{
			Type: 0,
			List: map[string]string{name: address},
		})
		if err != nil || resp1.StatusCode != 0 {
			log.Errorln("UpdateDispatcherList err:", err)
		}
	}
	log.Printf("register cluster manager, name:%v,state:%s", name, resp2.State)
	return nil
}
