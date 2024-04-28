package client

import (
	"context"
	"github.com/ServerlessOS/galaxy/constant"
	"github.com/ServerlessOS/galaxy/proto"
	"google.golang.org/appengine/log"
	"google.golang.org/grpc"
)

var (
	Sclient        = make(map[string]proto.SchedulerClient)
	SclientDefault proto.SchedulerClient //
	Dclient        = make(map[string]proto.DispatcherClient)
	DclientDefault proto.DispatcherClient
	Gclient        = make(map[string]proto.GatewayClient)
	GclientDefault proto.GatewayClient
	Nclient        = make(map[string]proto.NodeClient)
	NclientDefault proto.NodeClient
	Cclient        = make(map[string]proto.CoordinatorClient)
	CclientDefault proto.CoordinatorClient
	Fclient        = make(map[string]proto.FuncManagerClient)
	FclientDefault proto.FuncManagerClient
)

func DialFuncManagerClient(id string, address string) error {
	connFuncManager, err := grpc.Dial(address+":"+constant.FuncManagerPort, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := proto.NewFuncManagerClient(connFuncManager)
	Fclient[id] = client
	if FclientDefault == nil {
		FclientDefault = client
	}
	return nil
}
func DialCoordinatorClient(id string, address string) error {
	connCoordinator, err := grpc.Dial(address+":"+constant.CoordinatorPort, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := proto.NewCoordinatorClient(connCoordinator)
	Cclient[id] = client
	if CclientDefault == nil {
		CclientDefault = client
	}
	return nil
}
func DialSchedulerClient(id string, address string) error {
	connScheduler, err := grpc.Dial(address+":"+constant.SchedulerPort, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := proto.NewSchedulerClient(connScheduler)
	Sclient[id] = client
	if SclientDefault == nil {
		SclientDefault = client
	}
	return nil
}
func DialDispatcherClient(id string, address string) error {
	connDispatcher, err := grpc.Dial(address+":"+constant.DispatcherPort, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := proto.NewDispatcherClient(connDispatcher)
	Dclient[id] = client
	if DclientDefault == nil {
		DclientDefault = client
	}
	return nil
}

func DialGatewayClient(id string, address string) error {
	connGateway, err := grpc.Dial(address+":"+constant.GatewayRpcPort, grpc.WithInsecure())
	if err != nil {
		log.Errorf(context.Background(), "dial gateway err")
		return err
	}
	client := proto.NewGatewayClient(connGateway)
	Gclient[id] = client
	if GclientDefault == nil {
		GclientDefault = client
	}
	return nil
}
func DialNodeClient(id string, address string) error {
	connNode, err := grpc.Dial(address+":"+constant.NodePort, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := proto.NewNodeClient(connNode)
	Nclient[id] = client
	if NclientDefault == nil {
		NclientDefault = client
	}
	return nil
}

func GetSchedulerClient(s ...string) proto.SchedulerClient {
	if s == nil || len(s) > 1 {
		return SclientDefault
	}
	return Sclient[s[0]]
}

func GetSchedulerClientList() map[string]proto.SchedulerClient {
	return Sclient
}

func GetDispatcherClient(s ...string) proto.DispatcherClient {
	if s == nil || len(s) > 1 {
		return DclientDefault
	}
	return Dclient[s[0]]
}

func GetDispatcherClientList() map[string]proto.DispatcherClient {
	return Dclient
}

func GetGatewayClient(s ...string) proto.GatewayClient {
	if s == nil || len(s) > 1 {
		return GclientDefault
	}
	return Gclient[s[0]]
}

func GetGatewayClientList() map[string]proto.GatewayClient {
	return Gclient
}

func GetNodeClient(s ...string) proto.NodeClient {
	if s == nil || len(s) > 1 {
		return NclientDefault
	}
	return Nclient[s[0]]
}

func GetNodeClientList() map[string]proto.NodeClient {
	return Nclient
}

func GetCoordinatorClient(s ...string) proto.CoordinatorClient {
	if s == nil || len(s) > 1 {
		return CclientDefault
	}
	return Cclient[s[0]]
}

func GetCoordinatorClientList() map[string]proto.CoordinatorClient {
	return Cclient
}

func GetFuncManagerClient(s ...string) proto.FuncManagerClient {
	if s == nil || len(s) > 1 {
		return FclientDefault
	}
	return Fclient[s[0]]
}

func GetFuncManagerClientList() map[string]proto.FuncManagerClient {
	return Fclient
}
