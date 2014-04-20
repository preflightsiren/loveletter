package main

import ()

type Round struct {
	Active             bool
	Deck               *Deck
	Players            []*Player
	currentPlayerIndex int
}

func (r *Round) NumberOfPlayers() int   { return len(r.Players) }
func (r *Round) CurrentPlayer() *Player { return r.Players[r.currentPlayerIndex] }
func (r *Round) Init(players []*Player) *Round {
	r.Active = true
	r.Deck = NewDeck()
	r.Players = players
	r.currentPlayerIndex = 0
	for i := 0; i < len(players); i++ {
		_, card : = r.Deck.Draw()
		players[i].CurrentRound = r
		players[i].ReceiveCard(card)
	}
	return r
}

func NewRound(players []*Player) *Round { return new(Round).Init(players) }
