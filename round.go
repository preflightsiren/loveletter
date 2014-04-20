package main

import (
	"fmt"
)

type Round struct {
	Deck               *Deck
	Players            []*Player
	currentPlayerIndex int
	ActivityLog        []string
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
	r.Log("A new deck has been shuffled")
	r.Players = players
	r.currentPlayerIndex = 0
	for i := 0; i < len(players); i++ {
		_, card := r.Deck.Draw()
		players[i].CurrentRound = r
		players[i].ReceiveCard(card)
	}
	r.Log("A new round has started")
	return r
}
func (r *Round) nextPlayer() {
	r.currentPlayerIndex = (r.currentPlayerIndex + 1) % r.NumberOfPlayers()
	r.Log(fmt.Sprintf("It's now %s's turn", r.CurrentPlayer().Name))
	if r.Active() {
		r.DrawForCurrentPlayer()
	}
}
func (r *Round) Discard(player *Player, card Card) {
	r.Log(fmt.Sprintf("%s just discarded card: %s", player.Name, card.Name()))
	r.Log(fmt.Sprintf("it has the effect of %s", player.Name, card.Action()))
	r.Deck.Discard(card)
	r.nextPlayer()
}
func (r *Round) DrawForCurrentPlayer() {
	var err error
	var card Card
	r.Log(fmt.Sprintf("%s draws a new card", r.CurrentPlayer().Name))
	err, card = r.Deck.Draw()
	if err != nil {
		//end round.
	}
	r.CurrentPlayer().ReceiveCard(card)
}

func (r *Round) PrintLog() {
	for i := 0; i < len(r.ActivityLog); i++ {
		fmt.Println(r.ActivityLog[i])
		if i >= 10 {
			return
		}
	}
}

func NewRound(players []*Player) *Round { return new(Round).Init(players) }
func (r *Round) Log(msg string) {
	r.ActivityLog = append(r.ActivityLog, msg)
}
