package main

import (
	"fmt"
	"sort"
)

type ordArray struct {
	a      []int32
	nElems int
}

func (o *ordArray) size() int {
	return o.nElems
}

func (o *ordArray) find(searchKey int32) int {
	return o.recFind(searchKey, 0, o.nElems-1)
}

func (o *ordArray) recFind(searchKey int32, lowerbound int, upperBound int) int {
	var curIn int

	curIn = (lowerbound + upperBound) / 2
	if o.a[curIn] == searchKey {
		return curIn
	} else if lowerbound > upperBound {
		return o.nElems
	} else {
		if o.a[curIn] < searchKey {
			return o.recFind(searchKey, curIn+1, upperBound)
		} else {
			return o.recFind(searchKey, lowerbound, curIn-1)
		}
	}

}

func (o *ordArray) insert(value int32) {

	o.a = append(o.a, value)

	// sort.Slice available in Go 1.8 with use of anonymous function
	sort.Slice(o.a, func(i, j int) bool { return o.a[i] < o.a[j] })

	// increase the number of elements in the array
	o.nElems += 1
}

func (o *ordArray) display() {
	for i := 0; i < o.nElems; i++ {
		fmt.Print(o.a[i], " ")
	}
	fmt.Println(" ")
	fmt.Println("Number of elements in array: ", o.nElems)
}

func main() {
	// create an array with a size of 100
	var arr ordArray

	// Initialise nElems to a value of zero
	arr.nElems = 0

	// add elems to the array using the insertion method
	arr.insert(int32(5))
	arr.insert(int32(10))
	arr.insert(int32(25))
	arr.insert(int32(15))
	arr.insert(int32(9))
	arr.insert(int32(20))
	arr.insert(int32(207))
	arr.insert(int32(28))

	// view elems
	arr.display()

	var searchKey int32
	fmt.Println("Enter search key: ")
	fmt.Scan(&searchKey)
	if arr.find(searchKey) != arr.size() {
		fmt.Println("Found item: ", searchKey)
	} else {
		fmt.Println("Unable to locate item: ", searchKey)
	}

}
