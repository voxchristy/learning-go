package main

// defer only works within a func
// LIFO strategy

import (
	"fmt"
)

func main() {
	fmt.Println("Close the file.")
	defer fmt.Println("Open the file.")
	defer fmt.Println("Statement 1.")
	defer fmt.Println("Statement 2.")
	defer fmt.Println("Statement 3.")
	defer fmt.Println("Statement 4.")
	defer fmt.Println("Statement 5.")
	fmt.Println("undeferred statement.")


}
