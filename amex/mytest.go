package main

import (
	"fmt"
	"sort"
)

func main() {
	oneSlice := []int{3, 9}
	secondSlice := []int{}
	thirdSlice := []int{0, 5, 8}
	fourthSlice := []int{2, 6, 7, 8}
	fifthSlice := []int{9, 10}

	ans := combineSlices(oneSlice, secondSlice, thirdSlice, fourthSlice, fifthSlice)
	fmt.Println(ans)
}

func combineSlices(slices ...[]int) []int {
	var finalSlice []int

	// first slice
	for _, slice := range slices {
		for _, e := range slice {
			finalSlice = append(finalSlice, e)
		}
	}

	// // second slice
	// for _, e := range secondSlice {
	// 	finalSlice = append(finalSlice, e)
	// }

	// // third slice
	// for _, e := range thirdSlice {
	// 	finalSlice = append(finalSlice, e)
	// }

	// // fourth slice
	// for _, e := range fourthSlice {
	// 	finalSlice = append(finalSlice, e)
	// }

	// fmt.Println("final slice: ", finalSlice)

	// sort the final slice
	sort.Ints(finalSlice)

	return finalSlice
}
