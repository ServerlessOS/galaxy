package cmd

import (
	"context"
	gateway_rpc "gateway/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	Sclient        = make(map[string]gateway_rpc.SchedulerClient)
	SclientDefault gateway_rpc.SchedulerClient //
	Dclient        = make(map[string]gateway_rpc.DispatcherClient)
	DclientDefault gateway_rpc.DispatcherClient
	Gclient        = make(map[string]gateway_rpc.GatewayClient)
	GclientDefault gateway_rpc.GatewayClient
	Nclient        = make(map[string]gateway_rpc.NodeClient)
	NclientDefault gateway_rpc.NodeClient
)

func clientInit() error {
	connCoordinator, err := grpc.Dial(coordinatorAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second*3))
	if err != nil {
		log.Fatalln("coordinator connect err,addr:", coordinatorAddr, "err:", err)
	}
	dispatchers, err := gateway_rpc.NewCoordinatorClient(connCoordinator).GatewayRegister(context.Background(), &gateway_rpc.GatewayRegisterReq{
		GatewayName: gatewayName,
	})
	if err != nil {
		log.Fatalln("dispatcher connect err", err)
	}
	log.Println("coordinator connect success.")
	for _, dispatcher := range dispatchers.Dispatchers {
		err = DialDispatcherClient(dispatcher.Name, dispatcher.Address)
		if err != nil {
			log.Println("dial dispatcher err:", err)
			return err
		}
		log.Println("dispatcher ", dispatcher.Name, "connect success.")
	}
	return nil
}
func DialSchedulerClient(id string, address string) error {
	connScheduler, err := grpc.Dial(address+":"+SchedulerPort, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := gateway_rpc.NewSchedulerClient(connScheduler)
	Sclient[id] = client
	if SclientDefault == nil {
		SclientDefault = client
	}
	return nil
}
func DialDispatcherClient(id string, address string) error {
	connDispatcher, err := grpc.Dial(address+":"+DispatcherPort, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := gateway_rpc.NewDispatcherClient(connDispatcher)
	Dclient[id] = client
	if DclientDefault == nil {
		DclientDefault = client
	}
	return nil
}

func DialGatewayClient(address string, id string) {
	connGateway, _ := grpc.Dial(address, grpc.WithInsecure())
	client := gateway_rpc.NewGatewayClient(connGateway)
	Gclient[id] = client
}
func DialNodeClient(address string, id string) {
	connNode, _ := grpc.Dial(address, grpc.WithInsecure())
	client := gateway_rpc.NewNodeClient(connNode)
	Nclient[id] = client
}

func GetSchedulerClient(s ...string) gateway_rpc.SchedulerClient {
	if s == nil || len(s) > 1 {
		return SclientDefault
	}
	return Sclient[s[0]]
}
func GetSchedulerClientList() map[string]gateway_rpc.SchedulerClient {
	return Sclient
}
func GetDispatcherClient(s ...string) gateway_rpc.DispatcherClient {
	if s == nil || len(s) > 1 {
		return DclientDefault
	}
	return Dclient[s[0]]
}
func GetDispatcherClientList() map[string]gateway_rpc.DispatcherClient {
	return Dclient
}
func GetGatewayClient(...string) gateway_rpc.GatewayClient {
	return nil
}
func GetGatewayClientList() map[string]gateway_rpc.GatewayClient {
	return Gclient
}
func GetNodeClient(...string) gateway_rpc.GatewayClient {
	return nil
}
func GetNodeClientList() map[string]gateway_rpc.GatewayClient {
	return Gclient
}
