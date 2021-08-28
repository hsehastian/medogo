package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	// add new element to array
	cards = append(cards, "Six of Spades")

	// iterate cards then print the index and element
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
