package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	printCat(&wg)
	printDog(&wg)

}

func printCat(wg *sync.WaitGroup) {
	wg.Add(1)
	for i := 0; i < 5; i++ {
		fmt.Println("Cat")
	}
	wg.Done()
}

func printDog(wg *sync.WaitGroup) {
	wg.Add(1)
	for i := 0; i < 5; i++ {
		fmt.Println("Dog")
	}
	wg.Done()
}
