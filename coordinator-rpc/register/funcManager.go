package register

import (
	"context"
	pb "github.com/ServerlessOS/galaxy/proto"
	"log"
	"time"
)

type FuncManager struct {
}

func (g *FuncManager) Register(req *pb.RegisterReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	log.Printf("registerForK8s dispatcher, name:%v,state:%s", name, resp.State)
}
