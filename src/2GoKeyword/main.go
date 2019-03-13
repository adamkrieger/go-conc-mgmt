package main

import (
	"log"
	"time"
)

func main(){
	go sayHello()

	waitForever()
}

func sayHello(){
	time.Sleep(1 * time.Second)

	log.Println("Hello asynchronously!")
}

func waitForever() {
	//avoid deadlock
	for{
		time.Sleep(100 * time.Second)
	}
}

