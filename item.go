package mycache

import "time"

// CacheItem ...
type cacheItem struct {
	key         interface{}
	value       interface{}
	lifeSpan    time.Duration
	timer       *time.Timer
	createTime  time.Time
	changeTime  time.Time
	accessCount uint64
}

// NewCacheItem ...
func newCacheItem(value interface{}) *cacheItem {
	return &cacheItem{
		value:      value,
		createTime: time.Now(),
		changeTime: time.Now(),
	}
}
