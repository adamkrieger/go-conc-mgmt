package main

import (
	"fmt"
	"time"

	"github.com/adamkrieger/go-conc-mgmt/hockey"
	"github.com/adamkrieger/go-conc-mgmt/state"
	"github.com/fatih/color"
)

var (
	team            *hockey.Team
	sleepInterval   = 0 * time.Millisecond
	startupInterval = 0 * time.Millisecond
)

func main() {
	//slowDown()
	//slowWaaaayDown()

	team.Players[1] = hockey.NewPlayer(team, 1, true, color.FgBlue)
	team.Players[2] = hockey.NewPlayer(team, 2, false, color.FgCyan)
	team.Players[3] = hockey.NewPlayer(team, 3, false, color.FgGreen)

	fmt.Println("The game is about to begin.")
	fmt.Println("Number of Players: ", len(team.Players))

	time.Sleep(1 * time.Second)

	go team.Players[1].Ex1PassItAround(sleepInterval)
	time.Sleep(startupInterval)
	go team.Players[2].Ex1PassItAround(sleepInterval)
	time.Sleep(startupInterval)
	go team.Players[3].Ex1PassItAround(sleepInterval)

	time.Sleep(3 * time.Second)
	fmt.Println("Game Over. ", state.Passes, " passes completed.")
}

func slowDown() {
	sleepInterval = 30 * time.Millisecond
	startupInterval = 10 * time.Millisecond
}

func slowWaaaayDown() {
	sleepInterval = 90 * time.Millisecond
	startupInterval = 30 * time.Millisecond
}
