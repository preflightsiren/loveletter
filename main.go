package main

import (
	"fmt"
)

var currentRound *Round

func getPlayers() []*Player {
	var name string
	var players []*Player
	for i := 0; i < 2; i++ {
		fmt.Printf("Enter Player%d name:", i)
		_, err := fmt.Scanln(&name)
		if err != nil {
			fmt.Print(err)
			continue
		}
		players = append(players, NewPlayer(name))
	}
	return players
}

func main() {
	fmt.Println("Starting new game\n")
	players := getPlayers()
	currentRound = NewRound(players)
	fmt.Printf("there are %d players\n", currentRound.NumberOfPlayers())
	for i := 0; i < currentRound.NumberOfActivePlayers(); i++ {
		fmt.Printf("%s has %s in their hand\n", players[i].Name, players[i].Hand[0].Name())
	}
	fmt.Println("====Log====")
	currentRound.PrintLog()
}

// Order of play:
// Deal cards to all players.
// Current player takes second card
// Current player makes decision + Discards card
// Action is resolved
// Next player
