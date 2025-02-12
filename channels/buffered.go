// https://www.scaler.com/topics/golang/buffered-and-unbuffered-channel-in-golang/

package main

import "fmt"

func main() {
	// declare a string channel
	cs := make(chan string)

	// cs <- "hello"
	// cs <- "world"

	go func() {
		cs <- "hello"
	}()

	//go.dev/doc/effective_go#concurrency
	go func() {
		fmt.Println(<-cs)
	}() // parenthesis indicates that we must close the function

}
