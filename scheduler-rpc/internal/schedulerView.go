package internal

import (
	"sync"
)

//	type SchedulerInfo struct {
//		NodeName string
//		Address  string
//	}
type SchedulerView struct {
	items  []*SchedulerInfo
	lock   *sync.Mutex
	index  int
	l      int
	maxMem int64
}

func NewSchedulerView() *SchedulerView {
	return &SchedulerView{
		items:  make([]*SchedulerInfo, 0),
		lock:   &sync.Mutex{},
		index:  0,
		l:      0,
		maxMem: 0,
	}
}
func (sv *SchedulerView) GetLen() int {
	sv.lock.Lock()
	defer sv.lock.Unlock()
	return sv.l
}
func (sv *SchedulerView) Add(s *SchedulerInfo) {
	sv.lock.Lock()
	defer sv.lock.Unlock()
	s.MemLimit = -1
	sv.items = append(sv.items, s)
	sv.l++
}

func (sv *SchedulerView) Delete(s *SchedulerInfo) {
	sv.lock.Lock()
	defer sv.lock.Unlock()
	maxMem := int64(0)
	var result []*SchedulerInfo
	for _, info := range sv.items {
		if info.Address != s.Address {
			if info.MemLimit != -1 && info.MemLimit > maxMem {
				maxMem = info.MemLimit
			}
			result = append(result, info)
		}
	}
	sv.maxMem = maxMem
	sv.items = result
	sv.l--
}
func (sv *SchedulerView) GetSchedulerAddr(requireMem int64) string {
	sv.lock.Lock()
	defer sv.lock.Unlock()
	cnt := 0
	if sv.l == 0 {
		return ""
	}

	for cnt < sv.l {
		sv.index = (sv.index + 1) % len(sv.items)
		if sv.items[sv.index].MemLimit >= requireMem || sv.items[sv.index].MemLimit == -1 {
			return sv.items[sv.index].Address
		}
		cnt++
	}
	return ""
}
func (sv *SchedulerView) SetSchedulerLimit(addr string, newLimit int64) {
	sv.lock.Lock()
	defer sv.lock.Unlock()
	for _, info := range sv.items {
		if info.Address == addr {
			if info.MemLimit > newLimit {
				info.MemLimit = newLimit
			}
			return
		}
	}
}
