package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	players = make(map[int]*player)
)

type player struct {
	id      int
	hasPuck bool
}

func (thisPlayer *player) play() {
	assigned := ((thisPlayer.id + 1) % len(players)) + 1
	fmt.Println(strconv.Itoa(thisPlayer.id), "-", strconv.Itoa(assigned))
	for {
		time.Sleep(90 * time.Millisecond)

		if players[assigned].hasPuck {
			players[assigned].hasPuck = false
			thisPlayer.hasPuck = true
			fmt.Println(strconv.Itoa(assigned), " passed to ", strconv.Itoa(thisPlayer.id))
		}
	}
}

func main() {
	players[1] = &player{id: 1, hasPuck: true}
	players[2] = &player{id: 2}
	players[3] = &player{id: 3}

	go players[1].play()
	time.Sleep(30 * time.Millisecond)
	go players[2].play()
	time.Sleep(30 * time.Millisecond)
	go players[3].play()

	time.Sleep(5 * time.Second)
}
