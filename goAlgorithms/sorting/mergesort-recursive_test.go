package main

import (
	"fmt"
	"testing"
)

// func TestInsertionOfElems(t *testing.T) {
// 	var arr DArray
// 	elemsToInsert := []int32{64, 21, 33, 70, 12, 85, 44, 3, 12, 17, 0, 99, 108, 36, 71, 89, 82}

// 	expected := []int32{64, 21, 33, 70, 12, 85, 44, 3, 12, 17, 0, 99, 108, 36, 71, 89, 82}

// 	for _, e := range elemsToInsert {
// 		arr.insert(e)
// 	}

// 	arr.display()

// 	for i, e := range arr.theArray {
// 		fmt.Println("arr.theArray: ", i, "- ", e)
// 		fmt.Println("expected elem: ", i, "- ", expected[i])
// 		if e != expected[i] {
// 			t.Errorf("Incorrect. Got: %v, Want: %v", e, expected[i])
// 		}
// 	}

// }

func TestMergeSortAlgorithm(t *testing.T) {
	var valOne, valTwo string
	valOne = "hello"
	valTwo = "world"
	fmt.Println(valOne + " " + valTwo)
	var arr DArray
	elemsToInsert := []int32{64, 21, 33, 70, 12, 85, 44, 3, 12, 17, 0, 99, 108, 36, 71, 89, 82}

	for _, e := range elemsToInsert {
		arr.insert(e)
	}

	arr.display()

	arr.mergeSort()

}
