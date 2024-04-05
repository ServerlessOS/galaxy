package internal

import "sync"

type FuncView struct {
	items map[string][]*FuncInfo
	lock  *sync.Mutex
}

func NewFuncView() *FuncView {
	return &FuncView{
		items: make(map[string][]*FuncInfo),
		lock:  &sync.Mutex{},
	}
}

func (f *FuncView) Add(funcInfo *FuncInfo) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.items[funcInfo.FuncName] = append(f.items[funcInfo.FuncName], funcInfo)
}
func (f *FuncView) Delete(funcInfo *FuncInfo) {
	f.lock.Lock()
	defer f.lock.Unlock()
	var result []*FuncInfo
	funcInfos := f.items[funcInfo.FuncName]
	for _, info := range funcInfos {
		if info.Address != funcInfo.Address {
			result = append(result, info)
		}
	}
	f.items[funcInfo.FuncName] = result
}
func (f *FuncView) Dispatch(funcName string) string {
	f.lock.Lock()
	defer f.lock.Unlock()
	funcInfos := f.items[funcName]
	for _, v := range funcInfos {
		if v.State {
			v.State = false
			return v.Address
		}
	}
	return ""
}
