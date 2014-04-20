package main

import (
	"fmt"
)

type Deck struct {
	Active         bool
	availableCards []Card
	burntCard      interface{}
	discardedCards interface{}
}

func (d *Deck) NumberInDeck() int { return len(d.availableCards) }

func NewDeck() Deck {
	princess := Card{8, "Princess", "Discarding this card loses the game"}
	countess := Card{7, "Countess", "Must discard this card when held with a King or Prince"}
	king := Card{6, "King", "Trade hands with another player of your choice"}
	prince := Card{5, "Prince", "Choose any player (including yourself) to discard his or her hand and draw a new card"}
	handmaiden := Card{4, "Handmaiden", "Until your next turn, ignore all effects from other players cards"}
	baron := Card{3, "Baron", "You and another player secretly compare hands. The player with the lower value is out of the round"}
	priest := Card{2, "Priest", "Look at another players hand"}
	guard := Card{1, "Guard", "Name a non-Guard card and choose another player. If that player has that card, he or she is out of the round"}

	var availableCards []Card

	availableCards = append(availableCards, princess)
	availableCards = append(availableCards, countess)
	availableCards = append(availableCards, king)
	for i := 0; i < 2; i++ {
		availableCards = append(availableCards, prince)
	}
	for i := 0; i < 2; i++ {
		availableCards = append(availableCards, handmaiden)
	}
	for i := 0; i < 2; i++ {
		availableCards = append(availableCards, baron)
	}
	for i := 0; i < 2; i++ {
		availableCards = append(availableCards, priest)
	}
	for i := 0; i < 5; i++ {
		availableCards = append(availableCards, guard)
	}
	deck := Deck{true, availableCards, nil, nil}

	return deck
}

func (d *Deck) Describe() {
	if d.Active {
		fmt.Println("Deck is active")
	} else {
		fmt.Println("Deck is not active")
	}
	for _, card := range d.availableCards {
		card.Describe()
	}
}

func (d *Deck) Draw(numberOfCards int) []Card {
	var drawnCards []Card
	for i := 0; i < numberOfCards; i++ {
		index := (i + 1) % len(d.availableCards)
		drawnCards = append(drawnCards, d.availableCards[index])
		d.availableCards = append(d.availableCards[:index], d.availableCards[index+1:]...)
		//Random Number
		// Pop card from availableCards
		//return cards
	}
	return drawnCards
}
