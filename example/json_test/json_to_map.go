package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	jsonStr := `
		{
        "name_tile": "liuXX",
        "age_size": 12
    }
`
	var mapResult map[string]interface{}
	if err :=json.Unmarshal([]byte(jsonStr),&mapResult);err !=nil{
		fmt.Errorf("error",err)
	}
	fmt.Println(mapResult["name_tile"])
}
