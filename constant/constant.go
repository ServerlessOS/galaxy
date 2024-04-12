package constant

const (
	gatewayCpuRequest     = 1
	gatewayMemRequest     = 2048
	gatewayDefaultPort    = 10001
	MaxDispatcherCacheNum = 10000
	MaxWorkerCacheNum     = 10000

	CoordinatorPort = "16000"
	DispatcherPort  = "16444"
	SchedulerPort   = "16445"
	NodePort        = "16446"
	GatewayHttpPort = "16447"
	GatewayRpcPort  = "16448"
	FuncManagerPort = "16449"

	NodeCpu = 4
	NodeMem = 4 * 1024
)
