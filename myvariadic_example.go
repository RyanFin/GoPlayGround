package main

import (
	"fmt"
	"reflect"
)

// variadic functions allow the user to enter any amount of input params into a function
// the amount can vary accordingly

func main() {
	multipleNums(12, 3, 4, 5, 6, 7, 8, 9, 123, 12, 2, 3, 4)
	multipleNums(0, 1, -2, -90)
	multipleNums(1)

	// heterogeneous -- interesting
	heterogenFunction("hello", 'h', 1.0, 56, -86787)
}

func multipleNums(nums ...int) {
	fmt.Println(nums)
	for i, e := range nums {
		fmt.Println("elem: ", i, " is ", e)
	}

}

// heterogeneous variadic function using interface{} type
func heterogenFunction(in ...interface{}) {
	fmt.Println(in)
	for _, e := range in {
		fmt.Println("e: ", reflect.TypeOf(e))
	}
}
