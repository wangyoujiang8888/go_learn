package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getSource(ch chan int32) {
	ra, rb := rand.Int31(), rand.Int31n(3)+1
	fmt.Println(rb)
	time.Sleep(time.Duration(rb) * time.Second)
	ch <- ra
}
func main6() {
	rand.Seed(time.Now().UnixNano())
	c := make(chan int32, 5)
	startTime := time.Now()
	for i := 0; i < cap(c); i++ {
		go getSource(c)
	}
	min := <-c
	fmt.Println(time.Since(startTime))
	fmt.Println("min source", min)
}
