package main

import (
	"fmt"
	"time"

	"github.com/adamkrieger/go-conc-mgmt/hockey"
	"github.com/fatih/color"
)

var (
	team            *hockey.Team
	sleepInterval   = 90 * time.Millisecond
	startupInterval = 30 * time.Millisecond
)

func init() {
	team = hockey.NewTeam()
}

func main() {
	team.Players[1] = hockey.NewPlayer(team, 1, true, color.FgBlue)
	team.Players[2] = hockey.NewPlayer(team, 2, false, color.FgCyan)
	team.Players[3] = hockey.NewPlayer(team, 3, false, color.FgGreen)

	fmt.Println("The game is about to begin.")
	fmt.Println("Number of Players: ", len(team.Players))

	time.Sleep(2 * time.Second)

	go team.Players[1].Ex1PassItAround(sleepInterval)
	time.Sleep(startupInterval)
	go team.Players[2].Ex1PassItAround(sleepInterval)
	time.Sleep(startupInterval)
	go team.Players[3].Ex1PassItAround(sleepInterval)

	time.Sleep(3 * time.Second)
	fmt.Println("Game Over")
}
