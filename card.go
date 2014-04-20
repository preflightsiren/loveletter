package main

import (
	"fmt"
)

type Card struct {
	points int
	name   string
	action string
}

func (c *Card) Name() string   { return c.name }
func (c *Card) Points() int    { return c.points }
func (c *Card) Action() string { return c.action }
func (c *Card) Describe() {
	fmt.Printf("Card: %s is worth %d points, when played has \"%s\" effect\n", c.Name(), c.Points(), c.Action())
}
