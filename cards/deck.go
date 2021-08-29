package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// `string(bs)` is type convertion from `[]byte` to `string`
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// below is one of the way how we generate random number in Go everytime we run the program
	// here we use `time.Now().UnixNano()` as a seed for the rand.NewSource() because UnixNano return int64
	s := rand.NewSource(time.Now().UnixNano())
	// after manage to generate new source we assign it to rand.New()
	r := rand.New(s)

	for i := range d {
		// use `r` which is "instance"(not sure how we call it) of type Rand
		np := r.Intn(len(d) - 1)

		// swap the element
		// d[i] get assign with d[np]
		// d[np] get assign with d[i]
		d[i], d[np] = d[np], d[i]
	}
}
