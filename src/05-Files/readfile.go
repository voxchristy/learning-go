package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	
	filename := "./abc.txt"

	content, err := ioutil.ReadFile(filename)
	checkerror(err)

	// content is in bytes so convert to string
	res:= string(content)
	fmt.Println("Bytes from file: ", content)
	fmt.Println("content from file: ", res)


}


func checkerror(err error){
	if err != nil{
		panic(err)
	}
}