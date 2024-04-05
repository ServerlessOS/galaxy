package server

import (
	"container/list"
	"google.golang.org/grpc"
)

// LRUCache represents the LRU cache structure.
type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
}

// entry represents a key-value pair in the cache.
type conn struct {
	key   string
	value *grpc.ClientConn
}

// NewLRUCache initializes a new LRU cache with the given capacity.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

// Get retrieves the value from the cache for the given key.
func (lru *LRUCache) Get(key string) *grpc.ClientConn {
	if elem, exists := lru.cache[key]; exists {
		lru.list.MoveToFront(elem)
		return elem.Value.(*conn).value
	}
	return nil
}

// Put inserts a key-value pair into the cache.
func (lru *LRUCache) Put(key string, value *grpc.ClientConn) {
	if elem, exists := lru.cache[key]; exists {
		lru.list.MoveToFront(elem)
		elem.Value.(*conn).value = value
	} else {
		if len(lru.cache) >= lru.capacity {
			// Remove the least recently used element from the cache
			oldest := lru.list.Back()
			if oldest != nil {
				delete(lru.cache, oldest.Value.(*conn).key)
				oldest.Value.(*conn).value.Close()
				lru.list.Remove(oldest)
			}
		}

		// Add the new entry to the cache and the front of the list
		newElem := lru.list.PushFront(&conn{key, value})
		lru.cache[key] = newElem
	}
}
