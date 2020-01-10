package go_cache2

import (
	"runtime"
	"sync"
	"time"
)

type Item struct {
	Object     interface{}
	Expiration int64
}

const (
	// For use with functions that take an expiration time.
	NoExpiration time.Duration = -1
	// For use with functions that take an expiration time. Equivalent to
	// passing in the same expiration duration as was given to New() or
	// NewFrom() when the cache was created (e.g. 5 minutes.)
	DefaultExpiration time.Duration = 0
)

/**
判断是否过期
*/
func (item *Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

type Cache struct {
	*cache
	// If this is confusing, see the comment at the bottom of New()
}

type cache struct {
	defaultExpiration time.Duration
	items             map[string]Item
	mux               sync.RWMutex
	janitor           *janitor
}

func (c *cache) Set(k string,v interface{},d time.Duration)  {
	var e int64
	//等于默认等于创建cache 传入的过期时间
	if d == DefaultExpiration{
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	item := Item{
		Object:     v,
		Expiration: e, // 等于 0 或者大于 0
	}
	c.mux.Lock()
	c.items[k] = item
	c.mux.Unlock()
}

func (c *cache) Get(k string)(interface{},bool)  {
	c.mux.RLock()
	item,found :=c.items[k]
	c.mux.RUnlock()
	if !found{
		return nil,false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration{
			return nil,false
		}
	}
	return item.Object,true
}

func (c *Cache) Delete(k string)  {
	c.mux.Lock()
	delete(c.items,k)
	c.mux.Unlock()
}

func (c *cache) GetExpiredTime(k string) (time.Time,bool)  {
	c.mux.RLock()
	item,found :=c.items[k]
	c.mux.RUnlock()
	if found{
		return time.Unix(0,item.Expiration),true
	}
	return  time.Unix(0,0),false
}

func (c *cache) DeleteExpired(){
	c.mux.Lock()
	now := time.Now().UnixNano()
	for k, v := range c.items {
		if v.Expiration > 0 && now > v.Expiration{
			delete(c.items,k)
		}
	}
	c.mux.Unlock()
}

func (c *cache)Count() int  {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return  len(c.items)
}


type janitor struct {
	Interval time.Duration
	stop     chan bool
}

func (j *janitor) Run(c *cache) {
	ticker := time.NewTicker(j.Interval)
	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		case <-j.stop:
			ticker.Stop()
			return
		}
	}
}

func stopJanitor(c *Cache) {
	c.janitor.stop <- true
}

func runJanitor(c *cache, ci time.Duration) {
	j := janitor{
		Interval: ci,
		stop:     make(chan bool),
	}
	c.janitor = &j
	go j.Run(c)
}
/**
返回一个cache对象
*/
func newCache(d time.Duration, item map[string]Item) *cache {
	if d == 0 {
		d = -1
	}
	c := cache{
		defaultExpiration: d,
		items:             item,
	}
	return &c
}

func newCacheWithJanitor(d time.Duration, ci time.Duration, item map[string]Item) *Cache {
	c := newCache(d, item)
	C := &Cache{c}
	if ci > 0 {
		runJanitor(c, ci)
		runtime.SetFinalizer(C,stopJanitor)
	}
	return C
}

func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items := make(map[string]Item)
	return newCacheWithJanitor(defaultExpiration, cleanupInterval, items)
}
