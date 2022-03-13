package main

import (
	"fmt"
	"time"
)

func main1() {
	c := make(chan int)
	go func(ch chan<- int,x int) {
		time.Sleep(time.Second)
		ch<-x*x
	}(c,3)

	done := make(chan struct{})
	go func(ch <-chan int) {
		time.Sleep(time.Second)
		n := <-ch
		fmt.Println(n)
		done <- struct{}{}
	}(c)

	<-done
	fmt.Println("bytes")

}