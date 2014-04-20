package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	Active         bool
	availableCards []Card
	burntCard      Card
	discardedCards []Card
}

const NUMBER_OF_GUARDS = 5
const NUMBER_OF_PRIESTS = 2
const NUMBER_OF_BARONS = 2
const NUMBER_OF_HANDMAIDENS = 2
const NUMBER_OF_PRINCES = 2

func (d *Deck) NumberInDeck() int { return len(d.availableCards) }

func NewDeck() *Deck { return new(Deck).Init() }

func (d *Deck) Init() *Deck {
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
	for i := 0; i < NUMBER_OF_PRINCES; i++ {
		availableCards = append(availableCards, prince)
	}
	for i := 0; i < NUMBER_OF_HANDMAIDENS; i++ {
		availableCards = append(availableCards, handmaiden)
	}
	for i := 0; i < NUMBER_OF_BARONS; i++ {
		availableCards = append(availableCards, baron)
	}
	for i := 0; i < NUMBER_OF_PRIESTS; i++ {
		availableCards = append(availableCards, priest)
	}
	for i := 0; i < NUMBER_OF_GUARDS; i++ {
		availableCards = append(availableCards, guard)
	}
	d.availableCards = availableCards
	d.Active = true
	_, d.burntCard = d.Draw()
	d.discardedCards = []Card{}

	return d
}

func (d *Deck) Describe() {
	if d.Active {
		fmt.Println("Deck is active")
	} else {
		fmt.Println("Deck is not active")
		fmt.Printf("Burnt card was:")
		d.burntCard.Describe()
	}
	for _, card := range d.availableCards {
		card.Describe()
	}
}

func (d *Deck) Update() {
	d.Active = (len(d.availableCards) > 0)
}

func (d *Deck) Draw() (error, Card) {
	var drawnCard Card
	if d.Active {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		index := r.Intn(len(d.availableCards))
		drawnCard = d.availableCards[index]
		d.availableCards = append(d.availableCards[:index], d.availableCards[index+1:]...)
		d.Update()
	} else {
		err := errors.New("no more cards can be drawn from this deck")
		return err, Card{}
	}
	return nil, drawnCard
}
