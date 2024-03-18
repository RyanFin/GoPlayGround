package main

import "fmt"

func main() {
	// make an unbuffered channel
	ch := make(chan string)

	// write data to the channel in a separate anonymous goroutine
	go func() { ch <- "hello world" }()

	// output that data
	fmt.Println(<-ch)
}
