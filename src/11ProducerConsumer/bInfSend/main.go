package main

import (
	"fmt"
	"time"
)

func main() {
	msgChan := make(chan string)

	//One of many styles
	go produceMsgs(msgChan)
	go receiveMsgs(msgChan)

	_ = <-make(chan bool)
}

func produceMsgs(msgChan chan<- string) {
	for 1 == 1 {
		msgChan <- time.Now().String()
		time.Sleep(1 * time.Second)
	}
}

func receiveMsgs(msgChan <-chan string) {
	for 1 == 1 {
		msg := <-msgChan
		fmt.Println(msg)
	}
}
