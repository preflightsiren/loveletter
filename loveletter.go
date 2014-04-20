package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting new game\n")
	deck := NewDeck()
	deck.Active = false
	deck.Describe()
}
