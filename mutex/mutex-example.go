package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)

	// create the mutex to synchronise access to the state
	var mutex = &sync.Mutex{}

	var readOps uint64 //unsigned ints are either zero or non-negative
	// var writeOps uint64

	// read operation
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				// use mutex when accessing the state
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				// iterate the value of the read operations that have taken place with an atomic operation
				atomic.AddUint64(&readOps, 1) // iterate value by 1
				time.Sleep(time.Millisecond)
			}
		}()
	}

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)

	// Use mutex when accessing the state variable
	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()

}
