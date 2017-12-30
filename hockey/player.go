package hockey

import (
	"strconv"

	"github.com/fatih/color"
)

//Player -
type Player struct {
	ID        int
	Number    string
	HasPuck   bool
	Team      *Team
	Colorizer *color.Color
}

//NewPlayer -
func NewPlayer(team *Team, id int, hasPuck bool, textColor color.Attribute) *Player {
	newPlayer := &Player{
		Team:      team,
		ID:        id,
		Number:    strconv.Itoa(id),
		HasPuck:   hasPuck,
		Colorizer: color.New(textColor),
	}

	return newPlayer
}
