package main

import (
	"fmt"
	"time"
)

func main() {
	msgChan := make(chan string)
	cancelChan := make(chan interface{})

	go receiveMsgs(msgChan, cancelChan)
	go produceMsgs(msgChan, cancelChan)

	_ = <-make(chan bool)
}

func produceMsgs(msgChan chan<- string, cancelChan <-chan interface{}) {
	????
	for 1 == 1 {
		msgChan <- time.Now().String()
		time.Sleep(1 * time.Second)
	}

	close(msgChan)
}

func receiveMsgs(msgChan <-chan string, cancelChan chan<- interface{}) {
	msgs := 0
	canceled := false
	for 1 == 1 {
		msg, chanOk := <-msgChan
		msgs++

		if !chanOk {
			return
		} else {
			fmt.Println(msg)
		}

		//Cancel after 10 msgs
		if !canceled && msgs >= 10 {
			close(cancelChan)
			canceled = true
		}
	}
}
