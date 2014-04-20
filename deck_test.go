package main

import (
	"testing"
)

func TestNewDeckIsReadyForPlay(t *testing.T) {
	deck := NewDeck()
	if !deck.Active() {
		t.Fail()
	}
	return
}

func TestNumberOfCardsInDeck(t *testing.T) {
	princess := Card{8, "Princess", "Discarding this card loses the game"}
	countess := Card{7, "Countess", "Must discard this card when held with a King or Prince"}
	var availableCards []Card
	var burntCard Card
	var discardedCards []Card

	availableCards = append(availableCards, princess)
	availableCards = append(availableCards, countess)
	deck := Deck{availableCards, burntCard, discardedCards}

	if !(deck.NumberInDeck() == 2) {
		t.Error("After initialisation, Deck did not contain 2 cards.")
	}
}

func TestDrawingLastCardFromDeck(t *testing.T) {
	princess := Card{8, "Princess", "Discarding this card loses the game"}
	var availableCards []Card
	var burntCard Card
	var discardedCards []Card
	var drawnCard Card
	var err error
	availableCards = append(availableCards, princess)
	deck := Deck{availableCards, burntCard, discardedCards}

	if deck.NumberInDeck() != 1 {
		t.Fatalf("Deck should contain 1 card, but contains %d cards", deck.NumberInDeck())
	}

	err, drawnCard = deck.Draw()
	if err != nil {
		t.Fatal(err)
	}

	if drawnCard.Name() != "Princess" {
		t.Error("The wrong card was drawn from the deck")
	}

	if deck.NumberInDeck() != 0 {
		t.Errorf("Deck should be empty, but contains %d cards", deck.NumberInDeck())
		deck.Describe()
	}

	if deck.Active() != false {
		t.Errorf("Expected deck.Active to be false but was %t", deck.Active())
	}
}

func ExampleDescribeDeckF() {
	princess := Card{8, "Princess", "Discarding this card loses the game"}
	var availableCards []Card
	var burntCard Card
	var discardedCards []Card
	availableCards = append(availableCards, princess)
	deck := Deck{availableCards, burntCard, discardedCards}
	deck.Describe()

	// Output:
	// Deck is active
	// Card: Princess is worth 8 points, when played has "Discarding this card loses the game" effect
}
