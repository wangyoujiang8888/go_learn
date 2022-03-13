package main

import (
	"log"
	"time"
)

type T = struct{}

func work(id int, ready chan T, done chan T) {
	<-ready
	log.Print("Worker#", id, "开始工作")
	time.Sleep(time.Duration(id+1) * time.Second)
	done <- T{}
}
func main9() {

	ready, done := make(chan T), make(chan T)

	go work(0, ready, done)
	go work(1, ready, done)
	go work(2, ready, done)

	//ready <- T{}
	//ready <- T{}
	//ready <- T{}
	close(ready)

	<-done
	<-done
	<-done

}
