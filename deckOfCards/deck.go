package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type, called deck, which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

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
		// Option 1: Log the error and return a call to newDeck()
		// Option 2: Log the error and quit the program entirely *We are going with this option*
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",") // splits single string into a slice of type string
	return deck(s)                      // type conversion
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // Creates a new source to use as argument for `rand.New()`
	r := rand.New(source)                           // New random number generator

	for i := range d { // Notice that we did not pass in reference to the element, "card", we are iterating over
		newPosition := r.Intn(len(d) - 1) // len(d) == length of slice "d" //

		d[i], d[newPosition] = d[newPosition], d[i] // Swaps the elements at position i and position newPosition with each other
	}
}
