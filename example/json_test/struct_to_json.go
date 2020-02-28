package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string   `json:"name_tILE11"`
	Age  int         `json:"AGE_SIZE11"`
}


func main()  {
	p := People{
		Name: "river",
		Age:  34,
	}
	b,err := json.Marshal(&p)
	if err != nil{

	}
	fmt.Println(string(b))

}
