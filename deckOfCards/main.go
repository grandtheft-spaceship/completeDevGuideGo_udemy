package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "Ten of Hearts")

	// fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
