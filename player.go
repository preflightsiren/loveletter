package main

import (
	"errors"
)

type Player struct {
	Name         string
	Hand         []Card
	CurrentRound *Round
}

func NewPlayer(name string) *Player { return new(Player).Init(name) }
func (p *Player) Init(name string) *Player {
	p.Name = name
	p.Hand = []Card{}
	return p
}

func (p *Player) ReceiveCard(card Card) {
	p.Hand = append(p.Hand, card)
}

func (p *Player) Discard() error {
	if len(p.Hand) == 0 {
		return errors.New("Cannot discard a card from an empty hand")
	}
	index := len(p.Hand) - 1
	discardedCard := p.Hand[index]
	p.Hand = append(p.Hand[:index], p.Hand[index+1:]...)
	p.CurrentRound.Discard(p, discardedCard)
	return nil
}
