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
