package main

import (
	"testing"
)

func TestRoundCreation(t *testing.T) {
	players := []*Player{NewPlayer("Player1"), NewPlayer("Player2")}
	r := NewRound(players)

	if r.Active() != true {
		t.Errorf("Expected Round.Active to be true but was %t", r.Active())
	}
}

func TestNumberOfPlayers(t *testing.T) {
	players := []*Player{NewPlayer("Player1"), NewPlayer("Player2")}
	r := NewRound(players)
	if r.NumberOfPlayers() != 2 {
		t.Errorf("Round did not have the correct number of players. Expected 2, but got %d", r.NumberOfPlayers())
	}
}

func TestInitialisationOfRound(t *testing.T) {
	players := []*Player{NewPlayer("Player1")}
	p1 := players[0]
	r := NewRound(players)
	if p1.CurrentRound != r {
		t.Error("Expected player to be enrolled in the current round but was not.")
	}
	if len(p1.Hand) != 1 {
		t.Errorf("Expected player to have 1 card in their hand, but had %d cards", len(p1.Hand))
	}
}
