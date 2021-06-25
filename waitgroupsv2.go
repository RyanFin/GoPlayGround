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

	wg.Add(len(sentences))

	for i := 0; i < len(sentences); i++ {
		go speak(i, &wg, sentences[i])
	}

	wg.Wait()
}

// Use waitgroup pointer
func speak(id int, wg *sync.WaitGroup, s string) {
	defer wg.Done()

	fmt.Println("worker: ", id, " - ", s)

	time.Sleep(1 * time.Millisecond)

}
