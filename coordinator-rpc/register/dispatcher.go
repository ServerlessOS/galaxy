package register

import (
	"context"
	assignor "coordinator_rpc/RendezousHashing"
	"coordinator_rpc/client"
	pb "github.com/ServerlessOS/galaxy/proto"
	log "github.com/sirupsen/logrus"
	"time"
)

type Dispatcher struct {
}

func (d *Dispatcher) Register(req *pb.RegisterReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	name, address := req.Name, req.Address
	log.Println("new dispatcher,name:", name, "\naddress:", address)
	disp := &assignor.Dispatcher{
		Name: name,
		Addr: address,
		Hash: 0,
	}
	Rh.Dispatchers[disp.Name] = disp
	err := client.DialDispatcherClient(name, address)
	if err != nil {
		log.Errorln("dial dispatcher err,", err)
	}

	//给gateway同步dispatcher
	for s, _ := range Rh.Gateways {
		resp1, err := client.GetGatewayClient(s).UpdateDispatcherList(ctx, &pb.UpdateListReq{
			Type: 0,
			List: map[string]string{name: address},
		})
		if err != nil || resp1.StatusCode != 0 {
			log.Errorln("UpdateDispatcherList err:", err)
		}
	}

	//给新的dispatcher同步scheduler信息
	//若不需要同步
	if len(SchedulerList) == 0 {
		return nil
	}
	client := client.GetDispatcherClient(req.Name)
	resp2, err := client.UpdateSchedulerView(ctx, &pb.SchedulerViewUpdate{List: SchedulerList, Action: "ADD"})
	if err != nil {
		log.Errorln(err)
		return err
	}
	log.Printf("register dispatcher, name:%v,state:%s", name, resp2.State)
	return nil
}
