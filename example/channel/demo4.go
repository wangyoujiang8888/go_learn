package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longRequestTIme() <-chan int32{
	r := make(chan int32)
	go func() {
		time.Sleep(3*time.Second)
		r<- rand.Int31n(100)
	}()
	return r
}

func main4()  {
	a,b := longRequestTIme(),longRequestTIme()
	fmt.Println(<-a,<-b)
}
