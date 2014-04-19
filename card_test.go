package main

import (
	"testing"
)

func TestCardAccessors(t *testing.T) {
	c := Card{8, "Princess", "Discarding this card loses the game"}
	if c.Name() != "Princess" {
		t.Error("Card should have the name \"Princess\" ")
	}
	if c.Points() != 8 {
		t.Error("Test card should be worth 8 points")
	}
	if c.Action() != "Discarding this card loses the game" {
		t.Error("Test card action is wrong")
	}
}

func ExampleDescribeCardF() {
	c := Card{8, "Princess", "Discarding this card loses the game"}
	c.Describe()

	// Output:
	// Card: Princess is worth 8 points, when played has "Discarding this card loses the game" effect
}
