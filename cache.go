package mycache

import (
	"sync"
)

type cacheDB struct {
	mutex *sync.RWMutex
	table map[string]*cacheTable
}

var cache = &cacheDB{
	mutex: new(sync.RWMutex),
	table: make(map[string]*cacheTable),
}

func (c *cacheDB) getTable(tableName string) *cacheTable {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if t, exist := c.table[tableName]; exist {
		return t
	}
	return nil
}

func (c *cacheDB) addTable(tableName string) *cacheTable {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	t := newCacheTable(tableName)
	c.table[tableName] = t
	return t
}

// New ...
func New(tableName string) *cacheTable {
	if cacheTable := cache.getTable(tableName); cacheTable != nil {
		return cacheTable
	}
	return cache.addTable(tableName)
}
