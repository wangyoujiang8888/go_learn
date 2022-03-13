package main

import (
	"fmt"
	"time"
)

func main3()  {
	ball := make(chan string)
	kickBall := func(playerName string) {
		for  {
			fmt.Println(<-ball,"传球")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}

	go kickBall("张三")
	go kickBall("王五")
	go kickBall("李四")
	go kickBall("王又江")

	ball<-"裁判"
	var c chan int
	<-c

}
