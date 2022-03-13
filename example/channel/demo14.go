package main

import (
	"fmt"
	"time"
)

type Request interface{}

func handle(r Request) { fmt.Println(r.(int)) }

const RateLimitPeriod = time.Minute
const RateLimit = 200 // 任何一分钟内最多处理200个请求

func handleRequests(requests chan Request) {
	quotas := make(chan struct{}, RateLimit)
	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		for range tick.C {
			select {
			case quotas <- struct{}{}:
			default:
			}
		}
	}()

	for r := range requests {
		<-quotas
		go handle(r)
	}
}

func main() {
	requests := make(chan Request)
	go handleRequests(requests)
	// time.Sleep(time.Minute)
	for i := 0; i < 100; i++ {
		requests <- i
	}
	select {}
}
