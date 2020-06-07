package main

// Go runtime is statically linked to the application
// memory is allocated/deallocated automatically
// Diff btw make() and new()
// when you allocate a map using
// new()
// -we will get back memory address, indicating the loaction of map
// -but map has zero memory storage(no init)
// -if u try to set a value, it will throw panic error
// "panic: assignment to entry in nil map"

// make()
// -allocates memory and initialises memory

import (
	"fmt"
)

func main() {

	// var m map[string]int
	// m["key"] = 42
	// fmt.Println(m)

	m := make(map[string]int)
	m["key"] = 42
	fmt.Println(m)

}
