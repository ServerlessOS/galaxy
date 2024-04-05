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

type Dispatcher struct {
	RootModule
}

func (d *Dispatcher) Register() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	name, address := d.Pod.Name, d.getPodIP()
	disp := &assignor.Dispatcher{
		Name: name,
		Addr: address,
		Hash: 0,
	}
	server.Rh.Dispatchers[disp.Addr] = disp
	err := cmd.DialDispatcherClient(name, address+":"+constant.DispatcherPort)
	if err != nil {
		log.Fatalln("dial dispatcher err,", err)
	}
	//若不需要同步
	if len(SchedulerList) == 0 {
		return
	}
	//给新的dispatcher同步scheduler信息
	client := cmd.GetDispatcherClient(d.Pod.Name)
	resp, err := client.UpdateSchedulerView(ctx, &pb.SchedulerViewUpdate{List: SchedulerList, Action: "ADD"})
	if err != nil {
		log.Fatalln("UpdateSchedulerView err,", err)
	}
	log.Printf("register dispatcher, name:%v,state:%s", name, resp.State)
}
