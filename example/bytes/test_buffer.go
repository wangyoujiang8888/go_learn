

package main

import (
	"bytes"
	"fmt"
)

func main() {


	b := bytes.NewBuffer([]byte("hello"))
	//b:= bytes.NewBuffer(nil)

	c,_ := b.ReadByte()

	fmt.Println(string(c))
	fmt.Println(b.String())

	fmt.Println(b.Cap())

	b.WriteRune('中')

	b.WriteString("wordsimriverwang")



	b.Grow(100)

	fmt.Println(b.Cap())


}

