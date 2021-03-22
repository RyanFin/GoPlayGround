package main

import "fmt"

// iota enum values will increment
const (
	C1 = iota
	C2
	C3
)

const (
	A1 = iota + 1
	A2
	A3
)

func main() {
	fmt.Println("iota output: ")
	fmt.Println(C1, C2, C3)
	fmt.Println("iota output 2: ")
	fmt.Println(A1, A2, A3)

}
