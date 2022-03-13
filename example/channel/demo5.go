package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longRequestTImeR(r chan<- int32) {
	time.Sleep(3*time.Second)
	r<-rand.Int31n(100)
}


func main5()  {
	ra, rb := make(chan int32), make(chan int32)
	go longRequestTImeR(ra)
	go  longRequestTImeR(rb)

	fmt.Println(<-ra,<-rb)



}
