package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func main() {
	scaleChan := make(chan int)
	failureChan := make(chan string)

	processorMap := make(map[string]*processor)

	scale := 3

	go func() {
		flip := false
		for {
			if flip {
				scaleChan <- 6
			} else {
				scaleChan <- 4
			}
			flip = !flip
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		select {
		case scale = <-scaleChan:

		case key := <-failureChan:
			fmt.Println("removing failed processor: ", key)
			delete(processorMap, key)
		case <-time.After(1 * time.Second):
			currentDifferential := scale - len(processorMap)
			fmt.Println("checkin scale. current: ", len(processorMap), " min: ", scale)
			for i := 0; i < currentDifferential; i++ {
				key, newProcessor := startProcessor(failureChan)
				fmt.Println("adding processor: ", key)
				processorMap[key] = newProcessor
			}
		}
	}
}

type processor struct {
	key string
}

func startProcessor(failureChan chan<- string) (key string, handle *processor) {
	key = uuid.New().String()
	handle = &processor{
		key: key,
	}

	go handle.run(failureChan)

	return
}

func (proc *processor) run(failureChan chan<- string) {
	defer func() {
		if r := recover(); r != nil {
			failureChan <- proc.key
		}
	}()

	for {
		src := rand.NewSource(time.Now().UnixNano())
		rnd := rand.New(src).Float32() * 100
		
		if rnd > 98 {
			fmt.Println(proc.key, "panicking")
			panic("at the disco")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
