package main

import "fmt"

//Make a struct to represent a person.
//Person has a firstName, lastName

//Step1. Define the struct

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "Jim@gmail.com",
			zipCode: 94000,
		},
	}

	//OUTPUT: firstName "Jim"
	//Get the memory address of Jim and store it in jimPointer
	//On that pointer, update the name. Which should update the value int the address.
	jimPointer.updateName("jimmy")
	//OUTPUT: firstName: "Jimmy"
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	//update the firstName of the value at address pointerToPerson
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
