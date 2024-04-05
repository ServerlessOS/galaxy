package cmd

type workerInformation struct {
	requestId int64
	funcName  string
	addr      string
}
type dispatcherInformation struct {
	requestId int64
	addr      string
}

// 支持每个gateway只负责部分任务
var workerCache map[string]workerInformation
var dispatcherCache map[string]dispatcherInformation
