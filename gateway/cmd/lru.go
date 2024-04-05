package cmd

import "container/list"

type LRUCache struct {
	Cap       int
	M         map[int]*list.Element
	freqs     *list.List
	timeCount int
}

type LruNode struct {
	Key, Val, time int
}

func NewLru(capacity int) LRUCache {
	return LRUCache{
		Cap:       capacity,
		M:         make(map[int]*list.Element, capacity+1),
		freqs:     list.New(),
		timeCount: 0,
	}
}

func (this *LRUCache) Get(key int) (val int, isMiss bool) {
	if v, ok := this.M[key]; ok {
		v.Value = LruNode{
			Key:  v.Value.(LruNode).Key,
			Val:  v.Value.(LruNode).Val,
			time: this.timeCount,
		}
		this.timeCount++
		return v.Value.(LruNode).Val, false
	}
	return 0, true
}

func (this *LRUCache) Set(key int, value int) (isMiss, isReplace, writebackOccurs bool) {
	if _, ok := this.M[key]; ok {
		this.M[key].Value = LruNode{
			Key:  key,
			Val:  value,
			time: this.timeCount,
		}
		this.timeCount++
		return false, false, writebackOccurs
	} else {
		if this.freqs.Len() >= this.Cap {
			denode := this.getOldestVisitNode()
			if denode.Value.(LruNode).Val == 1 {
				writebackOccurs = true
			}
			delete(this.M, denode.Value.(LruNode).Key)
			denode.Value = LruNode{
				Key:  key,
				Val:  value,
				time: this.timeCount,
			}
			this.timeCount++
			this.M[key] = denode
			isReplace = true
		} else {
			node := this.freqs.PushFront(LruNode{
				Key:  key,
				Val:  value,
				time: this.timeCount,
			})
			this.timeCount++
			this.M[key] = node
		}
		return true, isReplace, writebackOccurs
	}
}
func (this *LRUCache) Keys() []int {
	keys := make([]int, len(this.M))
	i := 0
	for node := this.freqs.Back(); node != nil; node = node.Prev() {
		keys[i] = node.Value.(LruNode).Key
		i++
	}
	return keys
}

// Peek 查看是否存在，但不更新相对位置
func (this *LRUCache) Peek(key int) (isDirty, isExist bool) {
	if val, ok := this.M[key]; ok {
		if val.Value.(LruNode).Val == 1 {
			isDirty = true
		}
		return isDirty, true
	}
	return false, false
}
func (this *LRUCache) getOldestVisitNode() *list.Element {
	temp := this.freqs.Front()
	OldestNode := this.freqs.Front()
	tempTime := OldestNode.Value.(LruNode).time
	for temp != nil {
		if tempTime > temp.Value.(LruNode).time {
			OldestNode = temp
			tempTime = OldestNode.Value.(LruNode).time
		}
		temp = temp.Next()
	}
	return OldestNode
}
