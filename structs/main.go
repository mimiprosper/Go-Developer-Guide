package main

import "fmt"

/* struc = data structure. collection of properties that are related together.

structs, int, float, strings, and bool = value type. Use pointers to change things in a function

slice, maps, channels, pointers, functions = reference type. Don't use pointers
*/

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
	alex := person{
		firstName: "alex",
		lastName:  "okoro",
		// struct inside a struct (embedded struct)
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 12345,
		},
	}

	// pointer
	alex.upDateName("alexander")
	alex.print()

}

//pointer
func (pointerToPerson *person) upDateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// function taking a reciever (object of a class in OOP)
func (p person) print() {
	fmt.Printf("%+v", p)
}

