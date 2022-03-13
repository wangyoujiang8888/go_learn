package main

import (
	"fmt"
	"time"
)

func main8() {
	done := make(chan struct{})
	go func() {
		fmt.Println("Hello")
		time.Sleep(time.Second)
		<-done
	}()

	done <- struct{}{}
	fmt.Println("World")
}
