package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine 监控中...")
				time.Sleep(2*time.Second)
			}
		}
	}()

	time.Sleep(10*time.Second)
	fmt.Println("发送停止信号")
	stop <- true ;

}
