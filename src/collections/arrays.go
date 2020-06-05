package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println("Complete array: ", colors)
	fmt.Println("Color at first position:", colors[0])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers array: ", numbers)

	fmt.Println("No of colors: ", len(colors))
	fmt.Println("No of numbers: ", len(numbers))

}
