package main

import (
	"testing"
)

func TestRoundCreation(t *testing.T) {
	players := []*Player{NewPlayer("Player1"), NewPlayer("Player2")}
	r := NewRound(players)

	if r.Active != true {
		t.Errorf("Expected Round.Active to be true but was %b", r.Active)
	}
}

func TestNumberOfPlayers(t *testing.T) {
	players := []*Player{NewPlayer("Player1"), NewPlayer("Player2")}
	r := NewRound(players)
	if r.NumberOfPlayers() != 2 {
		t.Errorf("Round did not have the correct number of players. Expected 2, but got %d", r.NumberOfPlayers())
	}
}
