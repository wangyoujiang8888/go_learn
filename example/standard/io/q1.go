package main

import (
	"bytes"
	"os"
)

func main()  {
	file, err := os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bytes.NewReader([]byte("Go语言中文网"))
	reader.WriteTo(os.Stdout)

	//writer := bufio.NewWriter(os.Stdout)
	//writer.ReadFrom(file)
	//writer.Flush()
}