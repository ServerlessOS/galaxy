package cmd

import (
	"context"
	"encoding/json"
	gateway_rpc "gateway/proto"
	"net/http"
	"strconv"
	"sync"
)

var gatewayList = make(map[string]string)
var gatewayList_mutex sync.Mutex

func getGatewayList(w http.ResponseWriter, req *http.Request) {
	listString, err := json.Marshal(gatewayList)
	if err != nil {
		panic(err)
	}
	w.Write(listString)
}

// extensionGateway 接受http请求，扩容gateway
func extensionGateway(w http.ResponseWriter, req *http.Request) {
	//todo：扩容时主动传递一个gatewayid用于幂等性校验，目前使用任务id
	//1.通过dispatcher获取新gateway实例
	//2.更新新实例和旧实例的gatewayList信息
	ctx := context.Background()
	dClient := GetDispatcherClient()
	list := make([]*gateway_rpc.UserRequest, 1)
	list[0] = &gateway_rpc.UserRequest{
		RequestId:  0,
		FuncName:   "Gateway",
		RequireCpu: gatewayCpuRequest,
		RequireMem: gatewayMemRequest,
	}
	resp, err := dClient.Dispatch(context.Background(), &gateway_rpc.UserRequestList{List: list})
	if err != nil {
		panic(err)
	}
	//todo:暂且用requestId作为每个gateway的标识
	gatewayList_mutex.Lock()
	gatewayList[strconv.Itoa(int(resp.GetRequestId()))] = resp.GetDestination()
	gatewayList_mutex.Unlock()
	//本地list以同步完成，接着是同步其它gateway的列表
	for _, client := range GetGatewayClientList() {
		client.UpdateGatewayList(ctx, &gateway_rpc.UpdateGatewayListReq{
			Type:        0, // 0是append追加，具体应该参考proto文件
			GatewayList: map[string]string{strconv.Itoa(int(resp.GetRequestId())): resp.GetDestination()},
		})
	}
	//对新扩容的实例用覆盖
	gClient := GetGatewayClient(strconv.Itoa(int(resp.GetRequestId())))
	gClient.UpdateGatewayList(ctx, &gateway_rpc.UpdateGatewayListReq{
		Type:        2, // 2是VERRIDE
		GatewayList: gatewayList,
	})
	w.Write([]byte("success"))
	return
}
func checkGatewayIdIsExistence(id string) bool {
	if _, ok := gatewayList[id]; ok {
		return true
	} else {
		return false
	}
}
