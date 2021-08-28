package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	// add new element to array
	cards = append(cards, "Six of Spades")

	// call receiver from type "deck"
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
