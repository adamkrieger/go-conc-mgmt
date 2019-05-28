package main

var importantRecords = map[int]string{}

func main() {
	go process()
	go input()

	<-make(chan bool)
}

func process() {
	for {
		for i := 0; i < 10; i++ {
			delete(importantRecords, i)
		}
	}
}

func input() {
	for {
		for i := 0; i < 10; i++ {
			importantRecords[i] = "very important value"
		}
	}
}
