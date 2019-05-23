package main

import (
	"fmt"
	"time"
)

func main() {
	msgChan := make(chan string)

	go receiveMsgs(msgChan)
	go produceMsgs(msgChan)

	_ = <-make(chan bool)
}

func produceMsgs(msgChan chan<- string) {
	for i := 1; i <= 10; i++ {
		msgChan <- time.Now().String()
		time.Sleep(1 * time.Second)
	}

	close(msgChan)
}

func receiveMsgs(msgChan <-chan string) {
	for 1 == 1 {
		msg, chanOk := <-msgChan
		if !chanOk {
			return
		} else {
			fmt.Println(msg)
		}
	}
}
