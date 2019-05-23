package main

func main() {
	msgChan := make(chan string)

	msgChan <- "hello"
}
