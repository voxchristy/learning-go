package main

import (
	"fmt"
)

func main() {
	n1, l1 := Fullname("Grace", "Maria")
	fmt.Printf("Fullname: %v, Char length: %v \n\n", n1, l1)

	n2, l2 := FullNameNakedReturn("Gianna", "Celin")
	fmt.Printf("Fullname: %v, Char length: %v \n\n", n2, l2)
}

// Fullname exported func	
func Fullname(f, l string )(string, int){
	full := f + " " + l
	length:=len(full)
	return full, length
}

// FullNameNakedReturn Naked func
func FullNameNakedReturn(f, l string)(full string, length int){
	full = f + " " + l
	length=len(full)
	return
}