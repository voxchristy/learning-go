package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
}

func checkerror(err error){
	if err != nil{
		panic(err)
	}
}
