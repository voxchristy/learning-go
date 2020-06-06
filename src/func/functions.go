package main
// In Go there is no function/method overloading
// All functions here starts with a lower case initial character
// That makes them private to the current package ie they are not exported outside the package
// if initial char is uppercase, the funcion becomes public


import (
	"fmt"
)

func main() {
	dosomething()
	
	sum:=addvalues(23, 45)
	fmt.Println("Sum: ", sum)

	sum2:=addvalues2(1, 2)
	fmt.Println("Sum2: ", sum2)

	sum3:=addallvalues(4, 3, 9, 7)
	fmt.Println("Sum3: ", sum3)

}
func dosomething(){
	fmt.Println("Doing something")
}

func addvalues(v1 int, v2 int) int{
	return v1+v2
}

//to be more concise, u can declare data type once, if type is same
func addvalues2(v1, v2 int) int{
	return v1+v2
}

func addallvalues(values ...int) int{
	sum:=0
	for i:=range values{
		sum+=values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}
