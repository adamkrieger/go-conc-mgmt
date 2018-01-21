package hockey

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/adamkrieger/go-conc-mgmt/state"
)

var (
	puckMtx sync.Mutex
)

//Ex2PassItAround - Tries to take the puck from the assigned player
func (thisPlayer *Player) Ex2PassItAround() {
	teamSize := len(thisPlayer.Team.Players)
	assigned := ((thisPlayer.ID + 1) % teamSize) + 1
	teammates := thisPlayer.Team.Players
	passer := teammates[assigned]

	fmt.Println(strconv.Itoa(thisPlayer.ID), "-", strconv.Itoa(assigned))

	for {
		//time.Sleep(sleepInterval)

		puckMtx.Lock()
		thisPlayer.ex2CodeThatHasToRunLocked(passer)
		puckMtx.Unlock()
	}
}

func (thisPlayer *Player) ex2CodeThatHasToRunLocked(passer *Player) {
	if passer.HasPuck {
		passer.HasPuck = false
		thisPlayer.HasPuck = true

		passer.Colorizer.Print(passer.Number)
		fmt.Print(" passed to ")
		thisPlayer.Colorizer.Println(thisPlayer.Number)

		state.Passes++
	}
}
