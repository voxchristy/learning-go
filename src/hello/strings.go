package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicitly typed string"
	fmt.Printf("Value of Str1: %v; Type of str1: %T\n", str1, str1)

	str2 := "An explicitly typed string"
	fmt.Printf("Value of Str2: %v; Type of str2: %T\n", str2, str2)

	fmt.Println("To upper: ", strings.ToUpper(str1))
	fmt.Println("Title case: ", strings.Title(str1))

	lvalue := "hello"
	uvalue := "HELLO"

	fmt.Println("Is equal ?", (lvalue == uvalue))
	fmt.Println("Is equal non case sensitive ?", strings.EqualFold(lvalue, uvalue))

	fmt.Println("Contains string 'exp': ", strings.Contains(str1, "exp"))
	fmt.Println("Contains string 'exp': ", strings.Contains(str2, "exp"))

}
