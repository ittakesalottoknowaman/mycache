package mycache

import (
	"sync"

	"mycache/table"
)

type cacheDB struct {
	mutex *sync.RWMutex
	table map[string]*table.CacheTable
}

var cache = &cacheDB{
	mutex: new(sync.RWMutex),
	table: make(map[string]*table.CacheTable),
}

func (c *cacheDB) getTable(tableName string) *table.CacheTable {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if t, exist := c.table[tableName]; exist {
		return t
	}
	return nil
}

func (c *cacheDB) addTable(tableName string) *table.CacheTable {
	c.mutex.Lock()
	defer c.mutex.RUnlock()
	t := table.NewCacheTable(tableName)
	c.table[tableName] = t
	return t
}

func (c *cacheDB) deleteTable(tableName string) {
	c.mutex.Lock()
	defer c.mutex.RUnlock()
	delete(cache.table, tableName)
}

func New(tableName string) *table.CacheTable {
	if cacheTable := cache.getTable(tableName); cacheTable != nil {
		return cacheTable
	}
	return cache.addTable(tableName)
}
