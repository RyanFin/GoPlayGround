package main

import (
	"fmt"
	"testing"
)

func TestTwoSumsFunction(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(arr, target))

	// My expected result for this test
	expected := []int{0, 1}

	// run a test for each element in the slice to check whether function is running correctly
	for i, e := range twoSum(arr, target) {
		if e != expected[i] {
			t.Errorf("Error with twoSums() func... got %v, want %v", e, expected[i])
		}
	}
}

func TestTwoSumsFunctionv2(t *testing.T) {
	arr := []int{8, 30, 1, 6, 67, 13, 90}
	target := 21
	// My expected result for this test
	expected := []int{0, 5}

	// run a test for each element in the slice to check whether function is running correctly
	for i, e := range twoSum(arr, target) {
		if e != expected[i] {
			t.Errorf("Error with twoSums() func... got %v, want %v", e, expected[i])
		}
	}
}

func TestNoSummableInts(t *testing.T) {
	arr := []int{3, 6, 9, 12}
	target := 1

	fmt.Println("TestNoSummableInts: length:", len(twoSum(arr, target)))

	// I expect an empty array of results
	if len(twoSum(arr, target)) != 0 {
		t.Errorf("Error. There should be no results here. got len(%v), want len(%v)", twoSum(arr, target), 0)
	}
}
