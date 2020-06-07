package main

// no need to add the break

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("******************CASE 1********************")

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	result := ""
	// No break; (fallthrough) statement is needed like C/Java/C#
	switch dow {
	case 1:
		result = "Sunday"
		// fallthrough
	case 7:
		result = "Saturday"
	default:
		result = "Weekday"
	}
	fmt.Println(result)

	fmt.Println("********************CASE 2************")
	x := -42
	switch {
	case x < 0:
		result = "Negative"
	case x == 0:
		result = "Zero"
	default:
		result = "Positive"
	}
	fmt.Println(result)

}
