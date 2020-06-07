package main

import (
	"fmt"
)

func main() {
	var x float64 = 42
	var res string
	fmt.Println("The number x is: ", x)
	if x < 0 {
		res = "The number is less than 0"
	} else if x == 0 {
		res = "The number is equal to 0"
	} else {
		res = "The number is greater than 0"
	}
	fmt.Println(res)
	fmt.Println("Value of x is retained: ", x)
	fmt.Println("************************")

	// We have the option of adding initial statement, within the conditional block
	// Here: "k := 0;" is added
	fmt.Println("Adding initial statement in IF block")
	var op string
	if k := 0; k == 0 {
		op = "Zero"
	} else if k < 0 {
		op = "Negative"
	} else {
		op = "Positive"
	}
	fmt.Println(op)
	fmt.Println("Here, value of k is destroyed by GC outside if block")
	// Error : undefined: k
	// fmt.Println("Here, value of k is destroyed by GC outside if block: ", k)
}
