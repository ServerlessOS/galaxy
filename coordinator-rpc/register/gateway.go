package register

import (
	"context"
	assignor "coordinator_rpc/RendezousHashing"
	"coordinator_rpc/client"
	"coordinator_rpc/server"
	"github.com/ServerlessOS/galaxy/constant"
	pb "github.com/ServerlessOS/galaxy/proto"
	"log"
	"time"
)

type Gateway struct {
}

func (g *Gateway) Register(req *pb.RegisterReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	name, address := req.Name, req.Address
	gateway := &assignor.Gateway{
		Name: name,
		Addr: address,
	}
	server.Rh.Gateways[gateway.Name] = gateway
	err := client.DialGatewayClient(name, address+":"+constant.GatewayRpcPort)
	if err != nil {
		log.Println("dial gateway err,", err)
	}

	//向其它gateway同步list、向本gateway同步list、dispatcher、funcManager
	AllgatewayList := make(map[string]string)
	for n, v := range server.Rh.Gateways {
		if n == name {
			//自己不同步
			continue
		}
		AllgatewayList[v.Name] = v.Addr
		resp, err := client.GetGatewayClient(n).UpdateGatewayList(ctx, &pb.UpdateListReq{
			//    APPEND = 0;
			//    REDUCE = 1;
			//    OVERRIDE = 2;
			Type: 0,
			List: map[string]string{v.Name: v.Addr},
		})
		if resp.StatusCode != 0 || err != nil {
			log.Println("UpdateGatewayList err:", err)
		}
	}
	//本gateway同步
	resp, err := client.GetGatewayClient(name).UpdateGatewayList(ctx, &pb.UpdateListReq{
		//    APPEND = 0;
		//    REDUCE = 1;
		//    OVERRIDE = 2;
		Type: 0,
		List: AllgatewayList,
	})
	//dispatcher
	AllDispatcherList := make(map[string]string)
	for n, v := range server.Rh.Dispatchers {
		AllDispatcherList[v.Name] = v.Addr
	}
	resp, err := client.GetGatewayClient(name).UpdateDispatcherList(ctx, &pb.UpdateListReq{
		Type: 2,
		List: AllDispatcherList,
	})
	log.Printf("registerForK8s dispatcher, name:%v,state:%s", name, resp.State)
}
