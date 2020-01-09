package cache2

import (
	"log"
	"sync"
	"time"
)

type CacheTable struct {
	sync.RWMutex
	// The table's name.
	name string
	// All cached items.
	items map[interface{}]*CacheItem
	// Timer responsible for triggering cleanup.
	cleanupTimer *time.Timer
	// Current timer duration.
	cleanupInterval time.Duration
	// The logger used for this table.
	logger *log.Logger
}

func (table *CacheTable) expirationCheck(){
	table.Lock()
	if table.cleanupTimer != nil{
		table.cleanupTimer.Stop()
	}
	now := time.Now()
	smallestDuration := 0 * time.Second
	for key, item := range table.items{
		item.RLock()
		lifeSpan := item.lifeSpan
		accessedOn:=item.accessedOn
		item.RUnlock()
		if lifeSpan == 0 {
			continue
		}
		//已经过期的删除
		if now.Sub(accessedOn) >= lifeSpan {
			table.deleteInternal(key)
		}else {
			//找到最小的过期的时间
			if smallestDuration == 0 || lifeSpan - now.Sub(accessedOn) < smallestDuration {
				smallestDuration = lifeSpan - now.Sub(accessedOn)
			}
		}
	}
	if smallestDuration > 0 {
		table.cleanupTimer = time.AfterFunc(smallestDuration, func() {
			go table.expirationCheck()
		})
	}
	table.Unlock()
}

func (table *CacheTable) addInternal(item *CacheItem){
	table.items[item.key] = item
	//等于 0 代表没有还没有定时器 小于在运行的定时器 也不需要检查
	expDur := table.cleanupInterval
	table.Unlock()
	if item.lifeSpan > 0 && (expDur == 0 ||item.lifeSpan < expDur){
		table.expirationCheck()
	}
}

func (table *CacheTable) Value(key interface{}) (*CacheItem, error){
	table.RLock()
	r,ok := table.items[key]
	table.RUnlock()
	if ok{
		r.keeplive()
		return r,nil
	}
	return nil,ErrKeyNotFound
}

func (table *CacheTable) Add(name string,lifeSpan time.Duration,data interface{}) *CacheItem {
	table.Lock()
	item := NewCacheItem(name,lifeSpan,data)
	table.addInternal(item)
	return item
}

func (table *CacheTable) NotFoundAdd(key interface{}, lifeSpan time.Duration, data interface{}) bool {
	table.Lock()

	if _, ok := table.items[key]; ok {
		table.Unlock()
		return false
	}

	item := NewCacheItem(key, lifeSpan, data)
	table.addInternal(item)

	return true
}


func (table *CacheTable) deleteInternal(key interface{}) (*CacheItem, error){
	item,ok := table.items[key]
	if !ok {
		return nil,ErrKeyNotFound
	}
	delete(table.items,key)
	return  item,nil
}

func (table *CacheTable) Delete(key interface{})  {
	table.Lock()
	defer table.Unlock()
	table.deleteInternal(key)
}
