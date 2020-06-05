package main

import (
	"fmt"
)

func main() {

	//case 1
	var p *int
	if p != nil {
		fmt.Println("Value of pointer p: ", *p)
	} else {
		println("Pointer p is nil.")
	}

	//case 2
	var v int = 42
	p = &v
	if p != nil {
		fmt.Println("Value of pointer p: ", *p)
	} else {
		println("Pointer p is nil.")
	}

	//implicit
	var f1 float64 = 42.34
	ptr := &f1
	fmt.Println("Value of implicit ptr: ", *ptr)

	*ptr = *ptr / 10
	fmt.Println("Value of (*ptr / 31): ", *ptr)
	fmt.Println("Value of f1: ", f1)
}
