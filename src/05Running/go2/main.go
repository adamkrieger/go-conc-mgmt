package main

//12700 go
//12532 go1
//12440

import "time"

func main() {
	for i := 0; i < 1000000; i++ {
		go thread()
	}

	<-make(chan bool)
}

func thread() {
	for {
		time.Sleep(5 * time.Second)
	}
}
