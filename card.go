package main

import ()

type Card struct {
	points int
	name   string
	action string
}

func (c Card) Name() string   { return c.name }
func (c Card) Points() int    { return c.points }
func (c Card) Action() string { return c.action }
