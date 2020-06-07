package main

import (
	"fmt"
	"os"
	"errors"
)

func main() {
	
	f, err := os.Open("file.txt")
	
	if err == nil{
		fmt.Println(f)
	}else{
		fmt.Println(err.Error())
	}

	//custom errors pkg
	myerror:=errors.New("My custom error message")
	fmt.Println(myerror)

	//??
	attendance := map[string]bool{
		"Ann": true,
		"Mike": false}

	attended, ok := attendance["Ann"]
	if ok{
		fmt.Println("Ann attended? ", attended)
	}else{
		fmt.Println("No info")
	}
	
}
