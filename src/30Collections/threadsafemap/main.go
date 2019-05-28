package main

import "fmt"

type record struct {
	key   int
	value string
}

func main() {
	inputChan := make(chan *record)
	processChan := make(chan int)

	go process(processChan)
	go input(inputChan)
	go protect(inputChan, processChan)

	<-make(chan bool)
}

func process(processChan chan<- int) {
	for {
		for i := 0; i < 10; i++ {
			processChan <- i
		}
	}
}

func input(inputChan chan<- *record) {
	for {
		for i := 0; i < 10; i++ {
			inputChan <- &record{key: i, value: "very important value"}
		}
	}
}

func protect(inputChan <-chan *record, processChan <-chan int) {
	var importantRecords = map[int]string{}

	for {
		select {
		case inputRecord := <-inputChan:
			importantRecords[inputRecord.key] = inputRecord.value
		case processKey := <-processChan:
			if val, exists := importantRecords[processKey]; exists {
				fmt.Println(val)
			}
			delete(importantRecords, processKey)
		}
	}
}
