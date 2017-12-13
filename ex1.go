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
	id       int
	assigned int
	hasPuck  bool
}

func (thisPlayer *player) play() {
	for {
		//time.Sleep(90 * time.Millisecond)

		if players[thisPlayer.assigned].hasPuck {
			players[thisPlayer.assigned].hasPuck = false
			thisPlayer.hasPuck = true
			fmt.Println(strconv.Itoa(thisPlayer.assigned), " passed to ", strconv.Itoa(thisPlayer.id))
		}
	}
}

func main() {
	players[1] = &player{id: 1, assigned: 3, hasPuck: true}
	players[2] = &player{id: 2, assigned: 1}
	players[3] = &player{id: 3, assigned: 2}

	go players[1].play()
	//time.Sleep(30 * time.Millisecond)
	go players[2].play()
	//time.Sleep(30 * time.Millisecond)
	go players[3].play()

	time.Sleep(5 * time.Second)
}
