package main

import (
	"testing"
)

func TestEndToEnd(t *testing.T) {
	players := []*Player{NewPlayer("Player1"), NewPlayer("Player2")}
	r := NewRound(players)

	if r.Deck.Active() != true {
		t.Error("Expected the deck to be active, but it was not")
	}
	if r.Active() != true {
		t.Error("Expected the round to have active, but it was not.")
	}

	r.DrawForCurrentPlayer()
	for i := 0; i < 13; i++ {
		err := r.CurrentPlayer().Discard()
		if err != nil {
			t.Log(r.ActivityLog)
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
