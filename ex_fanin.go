package main

import (
	"fmt"
)

func main() {
	values := []int{
		34,
		54,
		25,
		93,
		100,
		9,
	}

	returnChan := make(chan string)

	for index, val := range values {
		go subtractFiveFromValueAsync(returnChan, index, val)
	}

	for i := 0; i < len(values); i++ {
		fmt.Println(<-returnChan)
	}
}

func subtractFiveFromValueAsync(returnChan chan<- string, index, value int) {
	result := value - 5
	returnChan <- fmt.Sprintf("Index %d Result %d", index, result)
}
