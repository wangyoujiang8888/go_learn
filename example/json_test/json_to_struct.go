package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name_tile"`
	Age  int    `json:"age_size"`
}

func main() {
	jsonStr := `
		{
        "name_tile": "liuXX",
        "age_size": 12
    }
`
	var people People
	json.Unmarshal([]byte(jsonStr),&people)
	fmt.Println(people)
}
