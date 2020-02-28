package slice

import (
	jsoniter "github.com/json-iterator/go"
)

func Struct2Map(s interface{}) map[string]interface{}  {
	m := make(map[string]interface{})
	b,_ :=jsoniter.Marshal(s)
	err := jsoniter.Unmarshal(b,&m)
	if err != nil{
		return nil
	}
	return m
}
