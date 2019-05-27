package main

import "fmt"

type result struct {
	success bool
	err     error
}

//Ex
// Data merged from two sources

func main() {
	resultChan := make(chan *result)

	//Async
	go doWorkAsync(resultChan)
	//Sync
	result := doWork()

	//Wait for Async to catch up
	resultAsync := <-resultChan

	//Merge results
	fmt.Println(result.success && resultAsync.success)
}

func doWork() *result {
	//Do some work

	return &result{
		success: true,
		err:     nil,
	}
}

func doWorkAsync(retChan chan<- *result) {
	//Do some work

	retChan <- &result{
		success: true,
		err:     nil,
	}
}
