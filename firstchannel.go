package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SendValue(c chan int) {
	// send value operator to send data to the channel
	c <- 8
}

func MyStringGoFunc(c chan string) {
	fmt.Println("Welcome to my string go func...")
	time.Sleep(1 + time.Second)
	c <- "Hello world Channel!"
}

func SendRandomInt(channel chan int) {
	fmt.Println("Generating random integer value...")
	channel <- rand.Intn(10)
}

// create a channel to communicate and send data between two goroutines
func main() {

	value := make(chan int, 3) // make channel
	defer close(value)         // close channel

	// call go routine
	go SendRandomInt(value)
	go SendRandomInt(value)
	go SendRandomInt(value)

	v := <-value

	fmt.Println(v)

}

func UnbufferedChannelExample() {
	// Example 1.
	values := make(chan int) // Unbuffered
	// prevent data from being sent to goroutines after channel is closed
	defer close(values)
	// call go routine
	go SendValue(values)

	// get int data from the channel
	v := <-values
	fmt.Println(v)
}

func BufferedChannelExample() {

	values := make(chan string, 2) //buffered channel with '2' lanes
	defer close(values)

	go MyStringGoFunc(values)
	go MyStringGoFunc(values)

	res := <-values

	fmt.Println("res: ", res)

}
