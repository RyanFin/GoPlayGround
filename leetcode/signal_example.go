package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// create a buffered channel for signals
	sigs := make(chan os.Signal, 1)
	// create a buffered channel for booleans
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// receive data from the sigs channel
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true // send true to the done channel
	}()

	fmt.Println("awaiting signal...")
	<-done
	fmt.Println("exiting.")
}
