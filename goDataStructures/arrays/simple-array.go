package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	sli := []int{2, 4, 6, 8, 10}

	for _, e := range sli {
		fmt.Print(e, " ")
	}
}
