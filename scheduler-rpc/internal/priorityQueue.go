package internal

import (
	"sync"
)

type PriorityQueue struct {
	items []*RequestItem
	len   uint64 //实际使用长度
	cap   uint64 //实际占用的空间的容量
	cmp   Comparator
	lock  sync.Mutex
	cond  *sync.Cond
}

func NewPriorityQueue() *PriorityQueue {

	RequestQueue := &PriorityQueue{
		items: make([]*RequestItem, 200),
		len:   0,
		cap:   200,
		lock:  sync.Mutex{},
	}
	RequestQueue.cond = sync.NewCond(&RequestQueue.lock)
	return RequestQueue
}
func (pq *PriorityQueue) Size() (num uint64) {
	if pq == nil {
		pq = NewPriorityQueue()
	}
	return pq.len
}

func (pq *PriorityQueue) Clear() {
	if pq == nil {
		pq = NewPriorityQueue()
	}
	pq.lock.Lock()
	//清空已分配的空间
	pq.items = make([]*RequestItem, 1)
	pq.len = 0
	pq.cap = 1
	pq.lock.Unlock()
}

func (pq *PriorityQueue) Empty() bool {
	if pq == nil {
		pq = NewPriorityQueue()
	}
	return pq.len == 0
}
func (pq *PriorityQueue) Push(item *RequestItem) {
	if pq == nil {
		pq = NewPriorityQueue()
	}
	pq.lock.Lock()
	defer pq.lock.Unlock()
	//判断是否存在比较器,不存在则寻找默认比较器,若仍不存在则直接结束
	if pq.cmp == nil {
		pq.cmp = GetCmp(item.Priority)
	}
	if pq.cmp == nil {
		return
	}
	//先判断是否需要扩容,同时使用和vector相同的扩容策略
	//即先翻倍扩容再固定扩容,随后在末尾插入元素e
	if pq.len < pq.cap {
		//还有冗余,直接添加
		pq.items[pq.len] = item
	} else {
		//冗余不足,需要扩容
		if pq.cap <= 65536 {
			//容量翻倍
			if pq.cap == 0 {
				pq.cap = 1
			}
			pq.cap *= 2
		} else {
			//容量增加2^16
			pq.cap += 65536
		}
		//复制扩容前的元素
		tmp := make([]*RequestItem, pq.cap, pq.cap)
		copy(tmp, pq.items)
		pq.items = tmp
		pq.items[pq.len] = item
	}
	pq.len++
	//到此时,元素以插入到末尾处,同时插入位的元素的下标为pq.len-1,随后将对该位置的元素进行上升
	//即通过比较它逻辑上的父结点进行上升
	pq.up(pq.len - 1)
	pq.cond.Broadcast()

}

func (pq *PriorityQueue) up(p uint64) {
	if p == 0 {
		//以及上升到顶部,直接结束即可
		return
	}
	if pq.cmp(pq.items[(p-1)/2].Priority, pq.items[p].Priority) > 0 {
		//判断该结点和其父结点的关系
		//满足给定的比较函数的关系则先交换该结点和父结点的数值,随后继续上升即可
		pq.items[p], pq.items[(p-1)/2] = pq.items[(p-1)/2], pq.items[p]
		pq.up((p - 1) / 2)
	}
}
func (pq *PriorityQueue) Pop() (e *RequestItem) {
	if pq == nil {
		pq = NewPriorityQueue()
	}
	pq.lock.Lock()
	if pq.Empty() {

		pq.cond.Wait()
	}

	//将最后一位移到首位,随后删除最后一位,即删除了首位,同时判断是否需要缩容
	e = pq.items[0]
	pq.items[0] = pq.items[pq.len-1]
	//pq.data[pq.len-1]=nil
	pq.len--
	//缩容判断,缩容策略同vector,即先固定缩容在折半缩容
	if pq.cap-pq.len >= 65536 {
		//容量和实际使用差值超过2^16时,容量直接减去2^16
		pq.cap -= 65536
		tmp := make([]*RequestItem, pq.cap, pq.cap)
		copy(tmp, pq.items)
		pq.items = tmp
	} else if pq.len*2 < pq.cap {
		//实际使用长度是容量的一半时,进行折半缩容
		pq.cap /= 2
		tmp := make([]*RequestItem, pq.cap, pq.cap)
		copy(tmp, pq.items)
		pq.items = tmp
	}
	//判断是否为空,为空则直接结束
	if pq.Empty() {
		pq.lock.Unlock()
		return
	}
	//对首位进行下降操作,即对比其逻辑上的左右结点判断是否应该下降,再递归该过程即可
	pq.down(0)
	pq.lock.Unlock()
	return e
}

func (pq *PriorityQueue) down(p uint64) {
	q := p
	//先判断其左结点是否在范围内,然后在判断左结点是否满足下降条件
	if 2*p+1 <= pq.len-1 && pq.cmp(pq.items[p].Priority, pq.items[2*p+1].Priority) > 0 {
		q = 2*p + 1
	}
	//在判断右结点是否在范围内,同时若判断右节点是否满足下降条件
	if 2*p+2 <= pq.len-1 && pq.cmp(pq.items[q].Priority, pq.items[2*p+2].Priority) > 0 {
		q = 2*p + 2
	}
	//根据上面两次判断,从最小一侧进行下降
	if p != q {
		//进行交互,递归下降
		pq.items[p], pq.items[q] = pq.items[q], pq.items[p]
		pq.down(q)
	}
}
func (pq *PriorityQueue) Top() (e *RequestItem) {
	if pq == nil {
		pq = NewPriorityQueue()
	}
	pq.lock.Lock()
	e = pq.items[0]
	pq.lock.Unlock()
	return e
}
