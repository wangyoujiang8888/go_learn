package main

import (
	"example/util/slice"
	"fmt"
)

func main()  {
	type People struct {
		Name string   `json:"name"`
		Age  int         `json:"age11"`
	}
	var p People
	p = People{
		Name: "river12",
		Age:  33,
	}
	out:= slice.Struct2Map(p)
	if out == nil{

	}
	fmt.Println(out)
}