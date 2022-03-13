package main

import "fmt"

func main2()  {
	c := make(chan int,2)
	c <-2
	c<-3
	close(c)
	fmt.Println(len(c),cap(c))
	x,ok := <-c
	fmt.Println(x,ok)
	fmt.Println(len(c),cap(c))
	x,ok = <-c
	fmt.Println(x,ok)
	fmt.Println(len(c),cap(c))

	x,ok = <-c
	fmt.Println(x,ok)
	fmt.Println(len(c),cap(c))
	//已关闭 接收数据
	x,ok = <-c
	fmt.Println(x,ok)
	fmt.Println(len(c),cap(c))
	//关闭写数据
	c<-7
}
