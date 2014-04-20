package main

import ()

type Round struct {
	Active        bool
	Deck          *Deck
	Players       []*Player
	CurrentPlayer *Player
}

func (r *Round) NumberOfPlayers() int { return len(r.Players) }
func (r *Round) Init(players []*Player) *Round {
	r.Active = true
	r.Deck = NewDeck()
	r.Players = players
	r.CurrentPlayer = r.Players[0]
	return r
}

func NewRound(players []*Player) *Round { return new(Round).Init(players) }
