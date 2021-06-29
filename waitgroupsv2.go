package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// 5 go routines
	sentences := []string{"hello world", "I am here", "To infinity and beyond", "well done", "Go Go Go"}

	for i := 0; i < len(sentences); i++ {
		// add only a single go routine to the wait group at a time
		wg.Add(1)
		// run the function
		go speak(i, &wg, sentences[i])
	}

	wg.Wait()
}

// Use waitgroup pointer for go routine function
func speak(id int, wg *sync.WaitGroup, s string) {
	// state that the goroutine is finished
	defer wg.Done()

	fmt.Println("worker: ", id, " - ", s)

	time.Sleep(1 * time.Millisecond)

}
