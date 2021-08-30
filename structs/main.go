package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// this is one of the way to declare type person
	// the down side for this approach is when someone accidendally swap the firsName and lastName order in type person
	// then the assigned value become wrong lastName = "Hery", firstName = "Sehastian"
	hery := person{"Hery", "Sehastian"}

	// this is the recommended way
	p := person{firstName: "Hery", lastName: "Sehastian"}

	fmt.Println(hery)
	fmt.Println(p)
}
