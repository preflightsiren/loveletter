package main

import ()

type Round struct {
	Deck               *Deck
	Players            []*Player
	currentPlayerIndex int
}

func (r *Round) NumberOfPlayers() int { return len(r.Players) }
func (r *Round) NumberOfActivePlayers() int {
	activePlayers := 0
	for i := 0; i < len(r.Players); i++ {
		if len(r.Players[i].Hand) > 0 {
			activePlayers = activePlayers + 1
		}
	}
	return activePlayers

}
func (r *Round) Active() bool {
	return r.Deck.Active() && r.NumberOfActivePlayers() >= 2
}
func (r *Round) CurrentPlayer() *Player { return r.Players[r.currentPlayerIndex] }
func (r *Round) Init(players []*Player) *Round {
	r.Deck = NewDeck()
	r.Players = players
	r.currentPlayerIndex = 0
	for i := 0; i < len(players); i++ {
		_, card := r.Deck.Draw()
		players[i].CurrentRound = r
		players[i].ReceiveCard(card)
	}
	return r
}
func (r *Round) nextPlayer() {
	r.currentPlayerIndex = (r.currentPlayerIndex + 1) % r.NumberOfPlayers()
}
func (r *Round) Discard(card Card) {
	r.Deck.Discard(card)
}
func (r *Round) DrawForCurrentPlayer() {
	var err error
	var card Card
	err, card = r.Deck.Draw()
	if err != nil {
		//end round.
	}
	r.CurrentPlayer().ReceiveCard(card)
}

func NewRound(players []*Player) *Round { return new(Round).Init(players) }
