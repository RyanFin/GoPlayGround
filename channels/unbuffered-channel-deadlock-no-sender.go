package main

import "fmt"

func main() {
	// create an unbuffered channel
	ch := make(chan string)

	// solution is to send some new data to the channel with a separate goroutine
	go func() {
		ch <- "hello world"
	}()

	// attempt to read without no sender to the channel will cause a deadlock!
	fmt.Println(<-ch)

}
