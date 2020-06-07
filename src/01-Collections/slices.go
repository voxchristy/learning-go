package main

import (
	"fmt"
	"sort"
)

func main() {
	// slice is an abstraction layer, that sits on top of an array
	// when u declare a slice,
	// runtime creates underlying array for you, allocates required memory and returns the requested slice
	// unlike arrays, slices are re-sizable,

	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	//append to slice
	colors = append(colors, "Purple")
	fmt.Println(colors)

	//remove first item from slice
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	//this syntax is same as below
	colors = append(colors[1:])
	fmt.Println(colors)

	//remove last item from slice
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	//create slice using, MAKE function: 3 args - type of slice items, initial length, capacity
	numbers := make([]int, 5, 5)
	numbers[0] = 12
	numbers[1] = 9
	numbers[2] = 534
	numbers[3] = 123
	numbers[4] = 1234
	fmt.Println(numbers)
	fmt.Println("Initial capacity: ", cap(numbers))

	//we can inspect slices capacity by cap function
	numbers = append(numbers, 255)
	fmt.Println(numbers)
	fmt.Println("capacity increased by init cap: ", cap(numbers))

	//sorting
	sort.Ints(numbers)
	fmt.Println("Sorted: ", numbers)

}
