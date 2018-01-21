package main

import (
	"fmt"
)

func main() {
	dataChannel := make(chan int)

	go runSomethingConcurrently(dataChannel)

	inlineResult := runSomethingOnThisThread()

	//If we put this line above the inlineResult, then it's sync again
	asyncResult := <-dataChannel

	fmt.Printf("Inline returned %v and Async returned %v \n", inlineResult, asyncResult)
}

func runSomethingConcurrently(returnChan chan<- int) {
	returnChan <- 27
	close(returnChan)
}

func runSomethingOnThisThread() int {
	return 89
}
