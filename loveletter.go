package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting new game\n")
	deck := NewDeck()
	deck.Describe()
}
