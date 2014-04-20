package main

import (
	"testing"
)

func TestNewDeckIsReadyForPlay(t *testing.T) {
	deck := NewDeck()
	if !deck.Active {
		t.Fail()
	}
	return
}

func TestNumberOfCardsInDeck(t *testing.T) {
	princess := Card{8, "Princess", "Discarding this card loses the game"}
	countess := Card{7, "Countess", "Must discard this card when held with a King or Prince"}
	var availableCards []Card

	availableCards = append(availableCards, princess)
	availableCards = append(availableCards, countess)
	deck := Deck{true, availableCards, nil, nil}

	if !(deck.NumberInDeck() == 2) {
		t.Error("After initialisation, Deck did not contain 2 cards.")
	}
}

func TestDrawingLastCardFromDeck(t *testing.T) {
	princess := Card{8, "Princess", "Discarding this card loses the game"}
	var availableCards []Card
	var drawnCards []Card
	availableCards = append(availableCards, princess)
	deck := Deck{true, availableCards, nil, nil}

	if deck.NumberInDeck() != 1 {
		t.Fatalf("Deck should contain 1 card, but contains %d cards", deck.NumberInDeck())
	}

	drawnCards = deck.Draw(1)
	if len(drawnCards) != 1 {
		t.Fatal("Draw failed to return a card")
	}
	if drawnCards[0].Name() != "Princess" {
		t.Error("The wrong card was drawn from the deck")
	}

	if deck.NumberInDeck() != 0 {
		t.Errorf("Deck should be empty, but contains %d cards", deck.NumberInDeck())
		deck.Describe()
	}
}

func TestDrawingCardsFromDeck(t *testing.T) {
	//var availableCards []Card
	//var drawnCards []Card
	deck := NewDeck()
	deck.Describe()

}

func ExampleDescribeDeckF() {
	princess := Card{8, "Princess", "Discarding this card loses the game"}
	var availableCards []Card
	availableCards = append(availableCards, princess)
	deck := Deck{true, availableCards, nil, nil}
	deck.Describe()

	// Output:
	// Deck is active
	// Card: Princess is worth 8 points, when played has "Discarding this card loses the game" effect
}
