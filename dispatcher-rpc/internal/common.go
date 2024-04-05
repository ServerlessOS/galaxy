package internal

type SchedulerInfo struct {
	NodeName string
	Address  string
	MemLimit int
}
type FuncInfo struct {
	FuncName  string
	Address   string
	Timestamp int64
	State     bool
}
