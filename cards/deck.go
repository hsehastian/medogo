package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// this is a receiver function, it can be assumed like instance method in PHP e.g. `$this->print()`
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
