/*
 * Simple caching library with expiration capabilities
 *     Copyright (c) 2013-2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE.txt
 */

package benchmark

import (
	"example/cache2"
	"example/go_cache2"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkNotFoundAdd(b *testing.B) {
	table := cache2.Cache("testNotFoundAdd")
	var finish sync.WaitGroup
	var added int32
	var idle int32

	fn := func(id int) {
		for i := 0; i < b.N; i++ {
			if table.NotFoundAdd(i, 0, i+id) {
				atomic.AddInt32(&added, 1)
			} else {
				atomic.AddInt32(&idle, 1)
			}
			time.Sleep(0)
		}
		finish.Done()
	}

	finish.Add(10)
	go fn(0x0000)
	go fn(0x1100)
	go fn(0x2200)
	go fn(0x3300)
	go fn(0x4400)
	go fn(0x5500)
	go fn(0x6600)
	go fn(0x7700)
	go fn(0x8800)
	go fn(0x9900)
	finish.Wait()
}


func BenchmarkGoCache2Add(b *testing.B){
	table := go_cache2.New(5*time.Minute,10*time.Minute)
	var finish sync.WaitGroup
	var added int32
	fn := func(id int) {
		for i := 0; i < b.N; i++ {
			 table.Set(string(i), i+id,go_cache2.DefaultExpiration)
			 atomic.AddInt32(&added, 1)
		}
		finish.Done()
	}
	finish.Add(10)
	go fn(0x0000)
	go fn(0x1100)
	go fn(0x2200)
	go fn(0x3300)
	go fn(0x4400)
	go fn(0x5500)
	go fn(0x6600)
	go fn(0x7700)
	go fn(0x8800)
	go fn(0x9900)
	finish.Wait()
}



