package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Ten of Hearts")

	// fmt.Println(cards)

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
