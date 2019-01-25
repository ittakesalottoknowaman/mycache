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
	t.items[key] = &item.CacheItem{
		Value: value,
	}
}

// SetWithExpire ...
func (t *CacheTable) SetWithExpire(key interface{}, value interface{}, expire time.Duration) {

}

// Expire ...
func (t *CacheTable) Expire(key interface{}, expire time.Duration) {

}

// Delete ...
func (t *CacheTable) Delete(key interface{}) {

}
