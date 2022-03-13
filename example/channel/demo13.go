package main

import (
	"fmt"
	"time"
)

func Tick(d time.Duration) chan struct{} {
	ch := make(chan struct{}, 1)
	go func() {
		for {
			time.Sleep(d)
			select {
			case ch <- struct{}{}:
			default:
			}
		}
	}()
	return ch
}

func main13() {
	d := time.Second
	now := time.Now()
	go func() {
		for range Tick(d) {
			fmt.Println("tick", time.Since(now))
		}
	}()
	go func() {
		for range time.Tick(time.Second) {
			fmt.Println("timeTick", time.Since(now))
		}
	}()

	select {}
}
