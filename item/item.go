package item

import "time"

// CacheItem ...
type CacheItem struct {
	key         interface{}
	value       interface{}
	LifeSpan    time.Duration
	Timer       *time.Timer
	accessCount uint64
}

// NewCacheItem ...
func NewCacheItem(value interface{}) *CacheItem {
	return &CacheItem{
		value: value,
	}
}

// Key ...
func (i *CacheItem) Key() interface{} {
	return i.Key
}

// Value ...
func (i *CacheItem) Value() interface{} {
	return i.value
}

// AccessCount ...
func (i *CacheItem) AccessCount() uint64 {
	return i.accessCount
}

// SetLifeSpan ...
func (i *CacheItem) SetLifeSpan(lifeSpan time.Duration) {

}
