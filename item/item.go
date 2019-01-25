package item

type CacheItem struct {
	value interface{}
}

func NewCacheItem(value interface{}) *CacheItem {
	return &CacheItem{
		value: value,
	}
}

func (i *CacheItem) Value() interface{} {
	return i.Value
}
