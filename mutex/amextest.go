package main

import (
	"fmt"
	"math/rand"
)

func Solution(N int) []int {
	// Implement your solution here
	fmt.Println(N)

	// random number from 1 - 1000

	var slice = make([]int, N)
	for i := 0; i < len(slice); i++ {
		// fmt.Println(randInt(1, 1000))

		rNum := randInt(1, 1000)
		slice[i] = rNum
	}

	fmt.Println(slice)

	fmt.Println(find_min(slice))

	return slice
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func find_min(A []int) int {
	var result int = 0
	for i := 1; i < len(A); i++ {
		if result > A[i] {
			result = A[i]
		}
	}
	return result
}

func main() {
	Solution(5)
}
