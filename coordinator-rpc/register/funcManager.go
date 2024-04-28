package register

import (
	"context"
	assignor "coordinator_rpc/RendezousHashing"
	"coordinator_rpc/client"
	"github.com/ServerlessOS/galaxy/constant"
	pb "github.com/ServerlessOS/galaxy/proto"
	"log"
	"time"
)

type FuncManager struct {
}

func (g *FuncManager) Register(req *pb.RegisterReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	name, address := req.Name, req.Address
	funcManager := &assignor.FuncManager{
		Name: name,
		Addr: address,
	}
	Rh.FuncManagers[funcManager.Name] = funcManager
	err := client.DialGatewayClient(name, address+":"+constant.FuncManagerPort)
	if err != nil {
		log.Println("dial funcManager err,", err)
	}
	//向所有gateway通告func-manager，此处没有让func-manager间交换数据，所以每一个func-manager需要保证可以处理任意函数的请求
	resp, err := client.GetGatewayClient(name).UpdateFuncManagerList(ctx, &pb.UpdateListReq{
		Type: 0,
		List: map[string]string{name: address},
	})
	if err != nil {
		return err
	}
	log.Printf("register funcManager, name:%v,state:%v \n", name, resp.StatusCode)
	return nil
}
