package main

func main() {
	initialPlayersPerSide := 5
	game := initGame(initialPlayersPerSide)

	game.play()

	waitChan := make(chan bool)
	_ = <-waitChan
}
