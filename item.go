package mycache

import (
	"time"
)

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

func (i *cacheItem) ttl() int64 {
	if i.lifeSpan == 0 {
		return 0
	}

	return (i.lifeSpan.Nanoseconds() - time.Now().Sub(i.changeTime).Nanoseconds()) / 1000000
}

func (i *cacheItem) isExpire() bool {
	if i.lifeSpan == 0 {
		return false
	}

	if time.Now().Sub(i.changeTime) > i.lifeSpan {
		return true
	}

	return false
}
