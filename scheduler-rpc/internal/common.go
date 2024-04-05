package internal

type SchedulerInfo struct {
	NodeName string
	Address  string
	MemLimit int64
}
type NodeResource struct {
	NodeName string
	HaveCpu  int64
	HaveMem  int64
	Address  string
	Port     string
}
type NodeResourceItem struct {
	Action string
	Node   NodeResource
}
type RequestInfo struct {
	RequestId      int64
	FunctionName   string
	RequireCpu     int64
	RequireMem     int64
	DispatcherAddr string
}
type RequestItem struct {
	Value    *RequestInfo // 实际的值
	Priority int64        // 优先级
}
