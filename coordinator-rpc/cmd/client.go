package cmd

import (
	pb "github.com/ServerlessOS/galaxy/proto"
	"time"

	"google.golang.org/grpc"
)

var (
	Sclient        = make(map[string]pb.SchedulerClient)
	SclientDefault pb.SchedulerClient //
	Dclient        = make(map[string]pb.DispatcherClient)
	DclientDefault pb.DispatcherClient
)

func DialSchedulerClient(id string, address string) error {
	connScheduler, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithTimeout(time.Second*3))
	if err != nil {
		return err
	}
	client := pb.NewSchedulerClient(connScheduler)
	Sclient[id] = client
	if SclientDefault == nil {
		SclientDefault = client
	}
	return nil
}
func DialDispatcherClient(id string, address string) error {
	connDispatcher, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithTimeout(time.Second*3))
	if err != nil {
		return err
	}
	client := pb.NewDispatcherClient(connDispatcher)
	Dclient[id] = client
	if DclientDefault == nil {
		DclientDefault = client
	}
	return nil
}

func GetSchedulerClient(s ...string) pb.SchedulerClient {
	if s == nil || len(s) > 1 {
		return SclientDefault
	}
	return Sclient[s[0]]
}
func GetSchedulerClientList() map[string]pb.SchedulerClient {
	return Sclient
}
func GetDispatcherClient(s ...string) pb.DispatcherClient {
	if s == nil || len(s) > 1 {
		return DclientDefault
	}
	return Dclient[s[0]]
}
func GetDispatcherClientList() map[string]pb.DispatcherClient {
	return Dclient
}
