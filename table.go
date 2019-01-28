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

	if item, exist := t.items[key]; exist {
		item.value = value
		item.changeTime = time.Now()
		return
	}
	item := newCacheItem(value)
	t.items[key] = item
}

// SetWithExpire ...
func (t *cacheTable) SetWithExpire(key interface{}, value interface{}, lifeSpan time.Duration) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if item, exist := t.items[key]; exist {
		item.value = value
	} else {
		t.Set(key, value)
	}

	t.Expire(key, lifeSpan)
}

// Get ...
func (t *cacheTable) Get(key interface{}) interface{} {
	t.mutex.RLock()

	item, exist := t.items[key]
	if !exist {
		return nil
	}

	// 判断是否超过过期时间
	if item.isExpire() {
		t.mutex.RUnlock()
		t.Delete(key)
		return nil
	}

	value := item.value
	t.mutex.RUnlock()
	return value
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
	// item.timer = time.AfterFunc(lifeSpan, func() { t.Delete(key) })
}

// TTL ...
func (t *cacheTable) TTL(key interface{}) int64 {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	item, exist := t.items[key]
	if !exist {
		return 0
	}

	ttl := item.ttl()
	return ttl
}
