package main

import (
	"fmt"
)

func main() {
	card := Card{8, "Princess", "Discarding this card loses the game"}
	//card.Points = 1
	fmt.Printf("Card: %s is worth %d points, when played has \"%s\" effect\n", card.Name(), card.Points(), card.Action())
	deck := NewDeck()
	deck.Describe()
}
