package main

import "fmt"

func main() {
	// buffered channels have a capacity/size value specified
	// they work on a FIFO basis
	ch := make(chan int, 3)

	// send data to the channel in the main go routine
	ch <- 5
	ch <- 10
	ch <- 15

	// receive data from the channel
	fmt.Println(<-ch) // 5 is the first out
	fmt.Println(<-ch) // 10 is the second out
	fmt.Println(<-ch) // 15 is the third out
}
