package main

import (
	"testing"
)

func TestEndToEnd(t *testing.T) {
	players := []*Player{NewPlayer("Player1"), NewPlayer("Player2")}
	p1 := players[0]
	r := NewRound(players)

	if r.Deck.Active() != true {
		t.Error("Expected the deck to be active, but it was not")
	}
	if r.Active() != true {
		t.Error("Expected the round to have active, but it was not.")
	}

	for i := 0; i < 14; i++ {
		r.DrawForCurrentPlayer()
		err := p1.Discard()
		if err != nil {
			t.Fatal(err)
		}
	}
	if r.Deck.Active() != false {
		t.Errorf("Expected the deck to be empty, but there was still %d cards left", r.Deck.NumberInDeck())
	}
	if r.Active() != false {
		t.Error("Expected the round to have finished")
	}
}
