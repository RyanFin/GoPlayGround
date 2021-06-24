package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		// send 'ping' to the channel
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}
