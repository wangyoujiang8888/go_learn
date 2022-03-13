package main

import (
	"fmt"
	"time"
)

func afterDuration(d time.Duration) <-chan struct{} {
	ch := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		ch <- struct{}{}
	}()
	return ch
}
func main10() {
	fmt.Println("Hello")
	<-afterDuration(1 * time.Second)
	fmt.Println("world")
	<-afterDuration(1 * time.Second)
	<-time.After(1 * time.Second)
	fmt.Println("time after ")

}
