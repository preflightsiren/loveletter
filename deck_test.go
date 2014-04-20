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

func TestDrawingCardsFromDeck(t *testing.T) {
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

func ExampleDescribeDeckF() {
	deck := NewDeck()
	deck.Describe()

	// Output:
	// Deck is active
	// Card: Princess is worth 8 points, when played has "Discarding this card loses the game" effect
	// Card: Countess is worth 7 points, when played has "Must discard this card when held with a King or Prince" effect
	// Card: King is worth 6 points, when played has "Trade hands with another player of your choice" effect
	// Card: Prince is worth 5 points, when played has "Choose any player (including yourself) to discard his or her hand and draw a new card" effect
	// Card: Prince is worth 5 points, when played has "Choose any player (including yourself) to discard his or her hand and draw a new card" effect
	// Card: Handmaiden is worth 4 points, when played has "Until your next turn, ignore all effects from other players cards" effect
	// Card: Handmaiden is worth 4 points, when played has "Until your next turn, ignore all effects from other players cards" effect
	// Card: Baron is worth 3 points, when played has "You and another player secretly compare hands. The player with the lower value is out of the round" effect
	// Card: Baron is worth 3 points, when played has "You and another player secretly compare hands. The player with the lower value is out of the round" effect
	// Card: Priest is worth 2 points, when played has "Look at another players hand" effect
	// Card: Priest is worth 2 points, when played has "Look at another players hand" effect
	// Card: Guard is worth 1 points, when played has "Name a non-Guard card and choose another player. If that player has that card, he or she is out of the round" effect
	// Card: Guard is worth 1 points, when played has "Name a non-Guard card and choose another player. If that player has that card, he or she is out of the round" effect
	// Card: Guard is worth 1 points, when played has "Name a non-Guard card and choose another player. If that player has that card, he or she is out of the round" effect
	// Card: Guard is worth 1 points, when played has "Name a non-Guard card and choose another player. If that player has that card, he or she is out of the round" effect
	// Card: Guard is worth 1 points, when played has "Name a non-Guard card and choose another player. If that player has that card, he or she is out of the round" effect

}
