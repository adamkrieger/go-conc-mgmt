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

	go doWorkAsync(resultChan)
	result := doWork()

	resultAsync := <-resultChan

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
