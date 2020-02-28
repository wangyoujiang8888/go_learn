package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "Liu xx"
	mapInstance["Age"] = 18
	mapInstance["Address"] = "广东 深圳"
	b,err := json.Marshal(&mapInstance)
	if err !=nil{

	}
	fmt.Println(string(b))
}
