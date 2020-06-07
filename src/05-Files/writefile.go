package main

import (
	"fmt"
	"io"
	"os"
	"io/ioutil"
)

func main() {
	
	content:="Content from Go, Hello world!!"

	// create file
	file, err:=os.Create("./createdfile.txt")
	checkerror(err)
	// defer will put the execution to wait till last 
	defer file.Close()

	// write content to created file
	length, err := io.WriteString(file, content)
	checkerror(err)

	fmt.Printf("All done with file of %v characters.", length)

	// using byte array, 0644 sets permission
	bytes := []byte(content)
	ioutil.WriteFile("./frombytes.txt", bytes, 0644)

}

func checkerror(err error){
	if err != nil{
		panic(err)
	}
}