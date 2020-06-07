package main

import (
	"fmt"
)

// Animal interface with 'Speak()' method
type Animal interface{
	Speak() string
}

// Dog struct 
type Dog struct{

}

// Speak func implements interface
func (d Dog) Speak() string {
	return "woof!"
}

// Cat struct
type Cat struct{

}

// Speak func implements interface
func (c Cat) Speak() string {
	return "meow!"
}

// Cow struct
type Cow struct{

}

// Speak func implements interface
func (w Cow) Speak() string {
	return "moo!"
}

func main() {
	poodle := Animal(Dog{})
	fmt.Println(poodle)

	animals:=[]Animal{Dog{}, Cat{}, Cow{}}

	for _, animal:=range animals{
		fmt.Println(animal.Speak())
	}
}
