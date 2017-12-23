package main

import (
	"fmt"
	"time"

	"github.com/adamkrieger/go-conc-mgmt/hockey"
)

var (
	team *hockey.Team
)

func init() {
	team = hockey.NewTeam()
}

func main() {
	team.Players[1] = hockey.NewPlayer(team, 1, true)
	team.Players[2] = hockey.NewPlayer(team, 2, false)
	team.Players[3] = hockey.NewPlayer(team, 3, false)

	fmt.Println("The game is about to begin.")
	fmt.Println("Number of Players: ", len(team.Players))

	time.Sleep(2 * time.Second)

	go team.Players[1].Play()
	time.Sleep(30 * time.Millisecond)
	go team.Players[2].Play()
	time.Sleep(30 * time.Millisecond)
	go team.Players[3].Play()

	time.Sleep(5 * time.Second)
	fmt.Println("Game Over")
}
