package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source() <-chan int32 {
	c := make(chan int32, 1) // 必须为一个缓冲通道
	go func() {
		ra, rb := rand.Int31(), rand.Intn(3)+1
		time.Sleep(time.Duration(rb) * time.Second)
		c <- ra
	}()
	return c
}

func main12() {
	rand.Seed(time.Now().UnixNano())
	var rnd int32
	// 阻塞在此直到某个数据源率先回应。
	select {
	case rnd = <-source():
	case rnd = <-source():
	case rnd = <-source():
	}
	fmt.Println(rnd)

	select {}
}
