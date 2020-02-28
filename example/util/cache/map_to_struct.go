package cache

import "github.com/mitchellh/mapstructure"

func Map2Struct(input map[string]string,out interface{}) error   {
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Metadata:         nil,
		Result:           out,
		TagName:          "json",
	}
	decoder,err :=mapstructure.NewDecoder(config)
	if err != nil{
		return  err
	}
	return decoder.Decode(input)
}
