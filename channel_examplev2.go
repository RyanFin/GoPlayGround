package main

import "fmt"

func main() {
	// make a messages unbuffered channel
	var messages = make(chan string)

	// send data to the channel using
	go func() {
		messages <- "this is my data payload"
	}()

	// Get data from the channel
	data := <-messages

	// output the data
	fmt.Println(data)

}
