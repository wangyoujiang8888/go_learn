package main

import "fmt"

func main11() {
	mutex := make(chan struct{}, 1)
	counter := 0
	increment := func() {
		mutex <- struct{}{}
		counter++
		<-mutex
	}
	done := make(chan struct{})

	increment1000 := func() {
		for i := 1; i <= 1000; i++ {
			increment()
		}
		done <- struct{}{}
	}

	go increment1000()

	go increment1000()

	<-done
	<-done

	fmt.Println(counter)
}
