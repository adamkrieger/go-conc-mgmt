package main

import (
	"fmt"
	"time"

	"github.com/adamkrieger/go-conc-mgmt/hockey"
	"github.com/adamkrieger/go-conc-mgmt/state"
	"github.com/fatih/color"
)

var (
	team *hockey.Team
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

	time.Sleep(1 * time.Second)

	go team.Players[1].Ex2PassItAround()
	go team.Players[2].Ex2PassItAround()
	go team.Players[3].Ex2PassItAround()

	time.Sleep(3 * time.Second)
	fmt.Println("Game Over. ", state.Passes, " passes completed.")
}
