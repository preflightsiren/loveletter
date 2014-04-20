package main

import (
	"testing"
)

func TestPlayerCreation(t *testing.T) {
	p := NewPlayer("Player1")
	if p.Name != "Player1" {
		t.Error("Player's name was not \"Player1\"")
	}
}

// Draw card

func TestDrawingAndDiscardingCard(t *testing.T) {
	princess := Card{8, "Princess", "Discarding this card loses the game"}

	players := []*Player{NewPlayer("Player1")}
	p := players[0]
	NewRound(players)
	p.ReceiveCard(princess)
	if len(p.Hand) != 2 {
		t.Errorf("Expected player to have 2 card in their hand, but they had %d", len(p.Hand))
	}
	err := p.Discard()
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Hand) != 1 {
		t.Errorf("Expected player to have 1 card in their hand, but they had %d", len(p.Hand))
	}
}

// func TestCardInHand(t *testing.T) {
// 	players := []*Player{NewPlayer("Player1")}
// 	p1 := players[0]
// 	r := NewRound(players)

// }
