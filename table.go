package mycache

import (
	"sync"
	"time"
)

type cacheTable struct {
	mutex     *sync.RWMutex
	tableName string
	items     map[interface{}]*cacheItem
}

func newCacheTable(tableName string) *cacheTable {
	return &cacheTable{
		mutex:     new(sync.RWMutex),
		tableName: tableName,
		items:     make(map[interface{}]*cacheItem),
	}
}

func (t *cacheTable) Set(key interface{}, value interface{}) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	var item *cacheItem
	var exist bool
	item, exist = t.items[key]
	if exist {
		item.value = value
	} else {
		item = newCacheItem(value)
		t.items[key] = item
	}

	item.changeTime = time.Now()
}

// SetWithExpire ...
func (t *cacheTable) SetWithExpire(key interface{}, value interface{}, lifeSpan time.Duration) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.items[key] = newCacheItem(value)
	t.Expire(key, lifeSpan)
}

// Get ...
func (t *cacheTable) Get(key interface{}) interface{} {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if item, exist := t.items[key]; exist {
		return item.value
	}

	return nil
}

// Delete ...
func (t *cacheTable) Delete(key interface{}) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if _, exist := t.items[key]; exist {
		delete(t.items, key)
		return true
	}
	return false
}

// Expire ...
func (t *cacheTable) Expire(key interface{}, lifeSpan time.Duration) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	item, exist := t.items[key]
	if !exist {
		return
	}

	item.changeTime = time.Now()
	item.lifeSpan = lifeSpan
	if item.timer != nil {
		item.timer.Reset(lifeSpan)
	}
	item.timer = time.AfterFunc(lifeSpan, func() { t.Delete(key) })
}

// TTL ...
func (t *cacheTable) TTL(key interface{}) int64 {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	item, exist := t.items[key]
	if !exist {
		return 0
	}

	if item.lifeSpan == 0 {
		return 0
	}

	return (item.lifeSpan.Nanoseconds() - time.Now().Sub(item.changeTime).Nanoseconds()) / 1000000
}
