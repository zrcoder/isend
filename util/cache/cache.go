package cache

import (
	"sync"
)

// a simple cache which can store key-values, thread safely.
type BytesCache struct {
	items    map[string][]byte
	rwMutex  *sync.RWMutex
	capacity int
}

func NewWithCapacity(capacity int) *BytesCache {
	c := &BytesCache{}
	c.items = make(map[string][]byte, capacity)
	c.rwMutex = new(sync.RWMutex)
	c.capacity = capacity
	return c
}

// add an item, if the key is aready exist, returns false directly.
func (p BytesCache) Add(key string, value []byte) (ok bool) {
	if _, found := p.Search(key); found {
		return false
	}
	p.rwMutex.Lock()
	if len(p.items) >= p.capacity {
		for k, _ := range p.items {
			delete(p.items, k)
			break
		}
	}
	p.items[key] = value
	p.rwMutex.Unlock()
	return true
}

// remove an item, if the key is not exist, returns false directly.
func (p BytesCache) Remove(key string) (ok bool) {
	if _, found := p.Search(key); !found {
		return false
	}
	p.rwMutex.Lock()
	delete(p.items, key)
	p.rwMutex.Unlock()
	return true
}

// replace an item, if the key is not exist, returns false directly.
func (p BytesCache) Replace(key string, newValue []byte) (ok bool) {
	if _, found := p.Search(key); !found {
		return false
	}
	p.rwMutex.Lock()
	p.items[key] = newValue
	p.rwMutex.Unlock()
	return true
}

func (p BytesCache) Search(key string) (value []byte, found bool) {
	p.rwMutex.RLock()
	value, found = p.items[key]
	p.rwMutex.RUnlock()
	return
}
