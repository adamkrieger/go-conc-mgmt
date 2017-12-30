package hockey

import (
	"fmt"
	"strconv"
	"time"

	"github.com/adamkrieger/go-conc-mgmt/state"
)

//Ex1PassItAround - Tries to take the puck from the assigned player
func (thisPlayer *Player) Ex1PassItAround(sleepInterval time.Duration) {
	teamSize := len(thisPlayer.Team.Players)
	assigned := ((thisPlayer.ID + 1) % teamSize) + 1
	fmt.Println(strconv.Itoa(thisPlayer.ID), "-", strconv.Itoa(assigned))
	for {
		time.Sleep(sleepInterval)

		teammates := thisPlayer.Team.Players

		passer := teammates[assigned]

		if passer.HasPuck {
			passer.HasPuck = false
			thisPlayer.HasPuck = true

			passer.Colorizer.Print(passer.Number)
			fmt.Print(" passed to ")
			thisPlayer.Colorizer.Println(thisPlayer.Number)

			state.Passes++
		}
	}
}
