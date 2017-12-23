package hockey

type Team struct {
	Players map[int]*Player
}

func NewTeam() *Team {
	team := &Team{
		Players: make(map[int]*Player),
	}
	return team
}
