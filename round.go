package main

import ()

type Round struct {
	Active        bool
	Deck          *Deck
	Players       []Player
	CurrentPlayer *Player
}

func (r *Round) NumberOfPlayers() int { return len(r.Players) }
func (r *Round) Init() *Round {
	r.Active = true
	r.Deck = NewDeck()
	r.Players = []Player{}
	r.CurrentPlayer = &r.Players[0]
	return r
}

func NewRound() *Round { return new(Round).Init() }
