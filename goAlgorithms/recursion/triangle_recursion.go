package main

import "fmt"

func triangle(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + triangle(n-1)
	}
}

func main() {
	var theNumber int
	fmt.Scan(&theNumber)

	ans := triangle(theNumber)
	fmt.Println("Triangle = ", ans)

}
