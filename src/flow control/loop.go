package main

import (
	"fmt"
)

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println("Final sum: ", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	fmt.Println("*************Using Range***************************")
	for i := range colors {
		fmt.Println(colors[i])
	}

	//No while loops, but we can mimic by using for loop
	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum: ", sum)
		if sum > 200 {
			goto eoplabel
		}
		if sum > 500 {
			break
		}
	}

eoplabel:
	fmt.Println("END OF PROGRAM")
}
