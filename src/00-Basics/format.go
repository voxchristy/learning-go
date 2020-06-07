package main

import (
	"fmt"
)

func main() {
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown fox."

	k, e := fmt.Println(str1, str2, str3)

	if e == nil {
		fmt.Println("String Length:", k)
	}
}
