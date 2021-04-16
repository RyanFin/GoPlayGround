package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	var theNumber int
	fmt.Scan(&theNumber)

	ans := factorial(theNumber)
	fmt.Println("Factorial: ", ans)
}
