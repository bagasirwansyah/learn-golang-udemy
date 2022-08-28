package main

import "fmt"

type deck []string

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
