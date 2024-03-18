package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Add two goroutines to the waitgroup
	wg.Add(2)

	// goroutine 1
	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	// goroutine 2
	go func() {
		i := 50
		for i < 100 {
			fmt.Println(i)
			i += 2
		}
		wg.Done()
	}()

	// Wait for both go routines in the waitgroup to finish
	wg.Wait()

	// Output that program is now fully complete
	fmt.Println("Program complete...")
}
