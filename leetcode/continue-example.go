package main

import "fmt"

func main() {
	a := 0

	for a < 25 {
		if a == 10 || a == 15 {
			a = a + 1
			fmt.Println("skipped value - ", a)
			continue
		}
		fmt.Println("a value: ", a)
		a++
	}
}
