package internal

import "sync"

type InstanceInfo struct {
	RequestId         int64
	FuncName          string
	Address           string
	DispatcherAddress string
}

type FIFO struct {
	items []*InstanceInfo
	lock  sync.Mutex
	cond  *sync.Cond
}

func NewFIFO() *FIFO {
	InstanceInformQueue := &FIFO{
		items: make([]*InstanceInfo, 0),
	}
	InstanceInformQueue.cond = sync.NewCond(&InstanceInformQueue.lock)
	return InstanceInformQueue
}

func (q *FIFO) Enqueue(item *InstanceInfo) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = append(q.items, item)
	q.cond.Signal()
}
func (q *FIFO) Dequeue() *InstanceInfo {
	q.lock.Lock()
	defer q.lock.Unlock()
	for len(q.items) == 0 {
		q.cond.Wait() // 阻塞等待直到队列非空
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *FIFO) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.items)
}
