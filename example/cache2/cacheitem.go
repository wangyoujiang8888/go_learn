package cache2

import (
	"sync"
	"time"
)

type CacheItem struct {
	sync.RWMutex

	// The item's key.
	key interface{}
	// The item's data.
	data interface{}
	// How long will the item live in the cache when not being accessed/kept alive.
	lifeSpan time.Duration

	// Creation timestamp.
	createdOn time.Time
	// Last access timestamp.
	accessedOn time.Time
	// How often the item was accessed.
	accessCount int64
}

func NewCacheItem(key interface{},lifeSpan time.Duration,data interface{}) *CacheItem  {
	var now = time.Now()
	return &CacheItem{
		key:         key,
		data:        data,
		lifeSpan:    lifeSpan,
		createdOn:   now,
		accessedOn:  now,
		accessCount: 0,
	}
}

func (item *CacheItem) keeplive()  {
	item.Lock()
	defer item.Unlock()
	item.accessedOn = time.Now()
	item.accessCount +=1
}

func (item *CacheItem) Data() interface{}  {
	return item.data
}




