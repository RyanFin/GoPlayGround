package main

import "fmt"

// variadic functions allow the user to enter any amount of input params into a function
// the amount can vary accordingly

func main() {
	multipleNums(12, 3, 4, 5, 6, 7, 8, 9, 123, 12, 2, 3, 4)
	multipleNums(0, 1, -2, -90)
	multipleNums(1)
}

func multipleNums(nums ...int) {
	fmt.Println(nums)
	for i, e := range nums {
		fmt.Println("elem: ", i, " is ", e)
	}

}
