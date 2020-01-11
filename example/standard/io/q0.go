package main

import (
	"fmt"
	"os"
)

func main()  {
	file, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余")
	n, err := file.WriteAt([]byte("Go语言中文网"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

}


