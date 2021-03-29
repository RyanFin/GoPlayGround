package main

import "fmt"

// Your aim is to create a recursive function that sorts a list of numbers
// Expected output: 1,4,6,9,12,39

// We're not testing your algorithmics, the aim here is to just convert the algorithm to code. :)

// You can test for syntax errors etc. by running your code at any time

// Algorithm:
// Pick a pivot (first element from the slice)
// Compare each other element to the pivot, put into two slices (bigger than pivot and smaller than pivot)
// Execute quicksort() on both the smaller slice and bigger slice (it's a recursive function)
// Return the concatenation of the smaller slice, and the bigger slice

func main() {
	fmt.Println(quicksort([]int{6, 9, 1, 4, 39, 12}))
}

func quicksort(list []int) []int {

	var ret []int
	//   get first element in slice
	pivot := list[0]

	var smaller []int
	var bigger []int

	for _, e := range list {
		if e > pivot {
			bigger = append(bigger, e)
		}
		if e < pivot {
			smaller = append(smaller, e)
		}
	}

	var sortedSmaller []int
	var sortedBigger []int

	if len(smaller) > 0 {
		sortedSmaller = quicksort(smaller)
	}

	if len(bigger) > 0 {
		sortedBigger = quicksort(bigger)
	}

	for _, e := range sortedSmaller {
		ret = append(ret, e)
	}

	ret = append(ret, pivot)

	for _, e := range sortedBigger {
		ret = append(ret, e)
	}

	return ret

}
