package hockey

//Team -
type Team struct {
	Players map[int]*Player
}

//NewTeam -
func NewTeam() *Team {
	team := &Team{
		Players: make(map[int]*Player),
	}
	return team
}
