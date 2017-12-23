package hockey

import (
	"fmt"
	"strconv"
	"time"
)

type Player struct {
	ID      int
	HasPuck bool
	Team    *Team
}

func NewPlayer(team *Team, id int, hasPuck bool) *Player {
	newPlayer := &Player{
		Team:    team,
		ID:      id,
		HasPuck: hasPuck,
	}

	return newPlayer
}

func (thisPlayer *Player) Play() {
	teamSize := len(thisPlayer.Team.Players)
	assigned := ((thisPlayer.ID + 1) % teamSize) + 1
	fmt.Println(strconv.Itoa(thisPlayer.ID), "-", strconv.Itoa(assigned))
	for {
		time.Sleep(1 * time.Nanosecond)

		teammates := thisPlayer.Team.Players
		if teammates[assigned].HasPuck {
			teammates[assigned].HasPuck = false
			thisPlayer.HasPuck = true
			fmt.Println(strconv.Itoa(assigned), " passed to ", strconv.Itoa(thisPlayer.ID))
		}
	}
}
