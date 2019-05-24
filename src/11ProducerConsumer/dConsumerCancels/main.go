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
	for 1 == 1 {
		select {
		case <-time.After(1 * time.Second):
			msgChan <- time.Now().String()
		case _, chanOK := <-cancelChan:
			if !chanOK {
				close(msgChan)
				return
			}
		}
	}
}

func receiveMsgs(msgChan <-chan string, cancelChan chan<- interface{}) {
	msgs := 0
	canceled := false
	for 1 == 1 {
		msg, chanOk := <-msgChan
		msgs++

		if !chanOk {
			return
		}

		fmt.Println(msg)

		//Cancel after 10 msgs
		if !canceled && msgs >= 10 {
			close(cancelChan)
			canceled = true
		}
	}
}
