package main

import (
	"fmt"
	"time"
)

// create method to execute
func world() {
	// print "world" 5 times
	for i := 0; i < 5; i++ {
		fmt.Println("world")
	}
}

func main() {
	// Run hello in the main Go Routine
	go world()
	fmt.Println("hello")

	go func(msg string) {
		fmt.Println(msg)
	}("after")

	// Goroutine scheduler needs to give up control from the main goroutine
	// The Sleep() method is required to do this
	time.Sleep(time.Second)
}
