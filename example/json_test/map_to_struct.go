package main

import (
	"example/util/cache"
	"fmt"
)


func main()  {
	type People struct {
		Name string `json:"name"`
		Age  int    `json:"age"`

	}
	mapInstance := make(map[string]interface{})
	mapInstance["name"] = "river"
	mapInstance["age"] = 28
	var people People
	cache.Map2Struct(mapInstance,&people)
	fmt.Println(people)
}
