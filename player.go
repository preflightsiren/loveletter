package main

import ()

type Player struct {
	Name string
	Hand []Card
}

func (p Player) DrawCardFromDeck() {
	card := Card{7, "Countess", "Must discard this card when held with a King or Prince"}
	p.Hand = append(p.Hand, card)
}

func NewPlayer(name string) *Player { return new(Player).Init(name) }
func (p *Player) Init(name string) *Player {
	p.Name = name
	p.Hand = []Card{}
	return p
}
