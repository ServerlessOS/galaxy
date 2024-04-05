package cache

import (
	"scheduler_rpc/internal"
	"sync"
)

type FIFO struct {
	items []*internal.NodeResourceItem
	lock  sync.Mutex
	cond  *sync.Cond
}

func NewFIFO() *FIFO {
	nodeResourceUpdateQueue := &FIFO{
		items: make([]*internal.NodeResourceItem, 0),
	}
	nodeResourceUpdateQueue.cond = sync.NewCond(&nodeResourceUpdateQueue.lock)
	return nodeResourceUpdateQueue
}
func (q *FIFO) Enqueue(item *internal.NodeResourceItem) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = append(q.items, item)
	q.cond.Signal()
}

// Dequeue 从队列头部取出一个元素
func (q *FIFO) Dequeue() *internal.NodeResourceItem {
	q.lock.Lock()
	defer q.lock.Unlock()
	for len(q.items) == 0 {
		q.cond.Wait() // 阻塞等待直到队列非空
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Len 返回队列的长度
func (q *FIFO) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.items)
}
