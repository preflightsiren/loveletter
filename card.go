package main

import (
	"fmt"
)

type Card struct {
	points int
	name   string
	action string
}

type Deck struct {
	numberInDeck   int
	availableCards []Card
	burntCard      interface{}
	discardedCards interface{}
}

func (c Card) Name() string   { return c.name }
func (c Card) Points() int    { return c.points }
func (c Card) Action() string { return c.action }
func (c Card) Describe() {
	fmt.Printf("Card: %s is worth %d points, when played has \"%s\" effect\n", c.Name(), c.Points(), c.Action())
}

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
	deck := Deck{13, availableCards, nil, nil}
	fmt.Printf("available cards has %d elements", len(availableCards))

	return deck
}

func (d Deck) Describe() {
	for _, card := range d.availableCards {
		card.Describe()
	}
}
