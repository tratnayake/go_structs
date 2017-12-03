package main

import (
	"fmt"
)

//Make a struct to represent a person.
//Person has a firstName, lastName

//Step1. Define the struct
type person struct {
	firstName string
	lastName  string
}

func main() {
	//Step2. Make an instance of that struct.

	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
