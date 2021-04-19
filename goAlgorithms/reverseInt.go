package main

import (
	"fmt"
)

func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		fmt.Println("remainder: ", remainder)
		new_int *= 10
		new_int += remainder
		fmt.Println("new_int: ", new_int)
		n /= 10
	}
	return new_int
}

func main() {
	fmt.Println("reverse result: ", reverse_int(123456))
}
