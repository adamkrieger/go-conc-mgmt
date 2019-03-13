package main

import (
	"fmt"
	"log"
	"time"
)

func main(){

	//Create a channel that sends/receives strings
	messageChannel := make(chan string)

	//Tell another goroutine to process messages
	go processMessages(messageChannel)

	msgs := []string{"hey", "there", "over", "channel", "bye!"}

	for idx, val := range msgs {
		messageChannel <- fmt.Sprintf("%d - %s", idx, val)
		time.Sleep(400 * time.Millisecond)
	}
}

// It's a good idea to set directions on parameter channels
func processMessages(messageChannel <-chan string) {
	for{
		msg := <-messageChannel

		log.Println(msg)
	}
}

