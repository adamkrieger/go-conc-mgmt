package main

import "fmt"

type taskData struct {
	resultChan chan *result
	input      string
}

type result struct {
	correlation int
	success     bool
	err         error
}

func main() {
	retChan := make(chan *result)

	tasks := map[int]*taskData{
		1: &taskData{resultChan: retChan, input: "1"},
		2: &taskData{resultChan: retChan, input: "2"},
		3: &taskData{resultChan: retChan, input: "3"},
	}

	for k, task := range tasks {
		go doWork(k, task)
	}

	for _ = range tasks {
		result := <-retChan
		fmt.Println("Task: ", result.correlation, " - ", result.success)
	}
}

func doWork(correlationID int, taskData *taskData) {

}
