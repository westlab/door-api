package common

import (
	"container/list"
)

type keyValue struct {
	key   string
	value interface{}
}

// LRUCache maintains cache upto the maxSize
//
// itemsList maintains order of element and data itself
// itemsMap holds instancce pointer of the data
// This implementation is ideal becaue only pointer is used for LRUCache
// ref:
//   https://github.com/dropbox/godropbox/blob/master/container/lrucache/lrucache.go
type LRUCache struct {
	itemsList *list.List
	itemsMap  map[string]*list.Element
	maxSize   int
}

// NewLRUCache creates a new LRUCache
func NewLRUCache(maxSize int) *LRUCache {
	if maxSize < 1 {
		panic("nonsensical LRU cache size specified")
	}

	return &LRUCache{
		itemsList: list.New(),
		itemsMap:  make(map[string]*list.Element),
		maxSize:   maxSize,
	}
}

// Set sets value to LRUCache
func (cache *LRUCache) Set(key string, val interface{}) {
	elem, ok := cache.itemsMap[key]
	if ok {
		// item already eixsts, so move it to the front of the list and update
		cache.itemsList.MoveToFront(elem)
		kv := elem.Value.(*keyValue)
		kv.value = val
	} else {
		elem = cache.itemsList.PushFront(&keyValue{key, val})
		cache.itemsMap[key] = elem

		// evict LRU entry if the cache is full
		if cache.itemsList.Len() > cache.maxSize {
			removedElem := cache.itemsList.Back()
			removedkv := removedElem.Value.(*keyValue)
			cache.itemsList.Remove(removedElem)
			delete(cache.itemsMap, removedkv.key)
		}
	}
}

// Get returns value from LRUCache and priority of cache is updated
func (cache *LRUCache) Get(key string) (val interface{}, ok bool) {
	elem, ok := cache.itemsMap[key]
	if !ok {
		return nil, false
	}

	// item exsits, so move it to front of list and return it
	cache.itemsList.MoveToFront(elem)
	kv := elem.Value.(*keyValue)
	return kv.value, true
}

// Len returns length of LRUCache
func (cache *LRUCache) Len() int {
	return cache.itemsList.Len()
}

// Delete deletes elements from LRUCache
func (cache *LRUCache) Delete(key string) (val interface{}, existed bool) {
	elem, existed := cache.itemsMap[key]

	if existed {
		cache.itemsList.Remove(elem)
		delete(cache.itemsMap, key)
	}

	return val, existed
}

// MaxSize returns maxSize in LRUCache
func (cache *LRUCache) MaxSize() int {
	return cache.maxSize
}
