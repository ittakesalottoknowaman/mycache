package table

import (
	"mycache/item"
	"sync"
	"time"
)

// CacheTable ...
type CacheTable struct {
	mutex     *sync.RWMutex
	tableName string
	items     map[interface{}]*item.CacheItem
}

// NewCacheTable ...
func NewCacheTable(tableName string) *CacheTable {
	return &CacheTable{
		mutex:     new(sync.RWMutex),
		tableName: tableName,
		items:     make(map[interface{}]*item.CacheItem),
	}
}

// Set ...
func (t *CacheTable) Set(key interface{}, value interface{}) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.items[key] = item.NewCacheItem(value)
}

// SetWithExpire ...
func (t *CacheTable) SetWithExpire(key interface{}, value interface{}, expire time.Duration) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
}

// Get ...
func (t *CacheTable) Get(key interface{}) interface{} {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if item, exist := t.items[key]; exist {
		return item.Value()
	}

	return nil
}

// Delete ...
func (t *CacheTable) Delete(key interface{}) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if _, exist := t.items[key]; exist {
		delete(t.items, key)
		return true
	}
	return false
}

// Expire ...
func (t *CacheTable) Expire(key interface{}, expire time.Duration) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
}
