package register

import (
	"context"
	assignor "coordinator_rpc/RendezousHashing"
	"coordinator_rpc/client"
	"fmt"
	pb "github.com/ServerlessOS/galaxy/proto"
	"google.golang.org/grpc/peer"
	"log"
	"net"
	"time"
)

type Gateway struct {
}

func (g *Gateway) Register(req *pb.RegisterReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	name, address := req.Name, req.Address
	log.Println("new gateway,address:", address, "\nname:", name)
	gateway := &assignor.Gateway{
		Name: name,
		Addr: address,
	}
	Rh.Gateways[gateway.Name] = gateway
	err := client.DialGatewayClient(name, address)
	if err != nil {
		pr, ok := peer.FromContext(ctx)
		if !ok {
			return fmt.Errorf("have some err")
		}
		// 获取客户端IP地址
		addr := pr.Addr
		tcpAddr, ok := addr.(*net.TCPAddr)
		if !ok {
			return fmt.Errorf("have some err")
		}
		IP := tcpAddr.IP.String()
		err = client.DialGatewayClient(name, IP)
		if err != nil {
			log.Println("dial gateway err,", err)
			return err
		}
	}
	//向其它gateway同步list、向本gateway同步list、dispatcher、funcManager
	AllgatewayList := make(map[string]string)
	for n, v := range Rh.Gateways {
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
		Type: 2,
		List: AllgatewayList,
	})
	if err != nil || resp.StatusCode != 0 {
		log.Println("UpdateGatewayList err:", err)
	}
	//dispatcher
	AllDispatcherList := make(map[string]string)
	for _, v := range Rh.Dispatchers {
		AllDispatcherList[v.Name] = v.Addr
	}
	resp, err = client.GetGatewayClient(name).UpdateDispatcherList(ctx, &pb.UpdateListReq{
		Type: 2,
		List: AllDispatcherList,
	})
	if err != nil || resp.StatusCode != 0 {
		log.Println("UpdateDispatcherList err:", err)
	}
	//funcManager
	AllFuncManagerList := make(map[string]string)
	for _, v := range Rh.FuncManagers {
		AllFuncManagerList[v.Name] = v.Addr
	}
	resp, err = client.GetGatewayClient(name).UpdateFuncManagerList(ctx, &pb.UpdateListReq{
		Type: 2,
		List: AllFuncManagerList,
	})
	if err != nil || resp.StatusCode != 0 {
		log.Println("UpdateFuncManagerList err:", err)
	}
	log.Printf("register gateway, name:%v", name)
	return nil
}
