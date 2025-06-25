package main

import (
	"fmt"
	"sync"
)

const (
	MAX = 200
)

// func evenNumbers() {
// 	for i := 0; i <= 200; i++ {
// 		if i%2 == 0 {
// 			fmt.Println("number: ", i)
// 		}
// 	}
// }

// func oddNumbers() {
// 	for i := 0; i <= 200; i++ {
// 		if i%2 == 1 {
// 			fmt.Println("number: ", i)
// 		}
// 	}
// }

// func main() {
// 	go oddNumbers()
// 	go evenNumbers()
// }

func printEven(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := 0; i <= MAX; i++ {
		if i%2 == 0 {
			// push even numbers to even channel
			ch <- i
		}
	}
	// close the channel
	close(ch)
}

func printOdd(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for i := 0; i <= MAX; i++ {
		if i%2 == 1 {
			// push odd numbers to odd channel
			ch <- i
		}
	}

	close(ch)
}

func main() {
	var wg sync.WaitGroup

	evenChan, oddChan := make(chan int), make(chan int)

	wg.Add(2)

	go printEven(&wg, evenChan)
	go printOdd(&wg, oddChan)

	go func() {
		for num := range evenChan {
			fmt.Println(num)
		}
	}()
	go func() {
		for num := range oddChan {
			fmt.Println(num)
		}
	}()

	// wait for completion of both tasks
	wg.Wait()
}
