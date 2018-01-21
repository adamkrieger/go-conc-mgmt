package hockey_f

type Team struct {
	players map[string]*Player
}

type Player struct {
	Receiving chan<- *Puck
	receiving <-chan *Puck
}

type Puck struct {
}
