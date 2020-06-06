package main

import (
	"fmt"
	"stringutil"
)

func main() {
	n1, l1 := stringutil.Fullname("Grace", "Maria")
	fmt.Printf("Fullname: %v, Char length: %v \n\n", n1, l1)

	n2, l2 := stringutil.FullNameNakedReturn("Gianna", "Celin")
	fmt.Printf("Fullname: %v, Char length: %v \n\n", n2, l2)
}