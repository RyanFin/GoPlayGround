package main

import (
	"fmt"
)

func main() {
	var n int
	var d int

	// enter the length of the array
	fmt.Scan(&n)
	fmt.Scan(&d)
	// enter the number of rotations
	//such as: 5 4

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		var in int
		fmt.Scan(&in) // enter each value entered into the array such as: 1 2 3 4 5
		arr[(i+n-d)%n] = in
	}

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

}
