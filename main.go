package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting new game\n")
	deck := NewDeck()
	for i := 0; i < 16; i++ {
		err, card := deck.Draw()
		if err != nil {
			fmt.Println(err)
		} else {
			card.Describe()
		}
	}
}
