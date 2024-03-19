package main

import "fmt"

func main() {
	// create a slice of ints spanning from 0 - 10
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var res string
	for _, e := range slice {
		res = evenOrOdd(e)
		fmt.Println(res)
	}

}

func evenOrOdd(e int) string {

	if e%2 == 0 {
		return fmt.Sprintf("%v - %s", e, "even")
	}

	if e%2 == 1 {
		return fmt.Sprintf("%v - %s", e, "odd")
	}

	return ""
}
