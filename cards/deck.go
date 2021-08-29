package main

import (
	"fmt"
	"strings"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// this is just a normal function
// the reason this function is not a "receiver" because to use "receiver" a variable need to be assigned with the `type`
// since the propose of this function is to generate list of cards so we dont need this as "receiver"
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// if there is variable that we don't use we can replace it with `_` (underscore)
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// this is a receiver function, it can be assumed like instance method in PHP e.g. `$this->print()`
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// this function return 2 values with type "deck"
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
