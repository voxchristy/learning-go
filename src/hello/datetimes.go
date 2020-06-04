package main

import (
	"fmt"
	"time"
)

func main() {

	//Date() and Now()
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)

	now := time.Now()
	fmt.Printf("Current time %s\n", now)

	//Month() Day() Weekday()
	fmt.Println("The month is ", now.Month())
	fmt.Println("The day is ", now.Day())
	fmt.Println("The weekday is ", now.Weekday())

	//Add date
	tomm := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v %v, %v\n", tomm.Weekday(), tomm.Month(), tomm.Day(), tomm.Year())

	//date formatting
	longFormat := "Monday, January 2, 2006"
	fmt.Println("Long format: ", tomm.Format(longFormat))

	shortFormat := "1/2/06"
	fmt.Println("Short format: ", tomm.Format(shortFormat))

}
