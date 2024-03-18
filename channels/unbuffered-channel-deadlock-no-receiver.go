package main

import "fmt"

func main() {
	// create an unbuffered channel
	ch := make(chan string)

	// deadlock should occur if there is no goroutine receiver for the sent data
	// e.g. printing out from the main goroutine only
	fmt.Println(<-ch)

	// solution is to create a new separate go routine to receive the data
	// go func() {
	// 	// send some data to the channel
	// 	ch <- "my data"
	// }()

	fmt.Println(<-ch)
}
