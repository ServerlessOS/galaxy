package server

import (
	"google.golang.org/grpc"
	"sync"
)

// LRUCache represents the LRU cache structure.
type ConnQueue struct {
	lock  sync.Mutex
	cond  *sync.Cond
	items []*grpc.ClientConn

	len int
}

// NewLRUCache initializes a new LRU cache with the given capacity.
func NewConnQueue() *ConnQueue {
	cq := &ConnQueue{
		lock: sync.Mutex{},

		items: make([]*grpc.ClientConn, 0),
		len:   0,
	}
	cq.cond = sync.NewCond(&cq.lock)
	return cq
}

// Get retrieves the value from the cache for the given key.
func (cq *ConnQueue) Dequeue() *grpc.ClientConn {
	cq.lock.Lock()
	defer cq.lock.Unlock()
	for cq.len == 0 {
		cq.cond.Wait() // 阻塞等待新元素进入队列
	}
	item := cq.items[0]
	cq.items = cq.items[1:]
	cq.len--
	return item

}
func (cq *ConnQueue) Enqueue(conn *grpc.ClientConn) {
	cq.lock.Lock()
	defer cq.lock.Unlock()
	cq.items = append(cq.items, conn)
	cq.len++
	cq.cond.Signal()
}
