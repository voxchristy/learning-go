package main

import (
	"fmt"
)

// Dog struct
type Dog struct{
	Breed string
	Weight int
	Sound string
}


// Speak func
func (d Dog) Speak(){
	fmt.Println(d.Sound)
}

// SpeakThreeTimes func 
func (d *Dog) SpeakThreeTimes(){
	d.Sound=fmt.Sprintf("%v! %v! %v! ", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {

	poodle:=Dog{"Poodle", 34, "woof"}
	fmt.Println(poodle)
	poodle.Speak()

	poodle.Sound="Arf"
	poodle.Speak()

	poodle.SpeakThreeTimes()
}
