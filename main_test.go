package main

import (
	"testing"
)

func TestEndToEnd(t *testing.T) {
	players := []*Player{NewPlayer("Player1")}
	p1 := players[0]
	r := NewRound(players)
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
