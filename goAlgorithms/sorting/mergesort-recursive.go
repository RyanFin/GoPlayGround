package main

import "fmt"

type DArray struct {
	theArray []int32
	nElems   int // length of array
}

func (d *DArray) insert(value int32) {
	// Insert value at index
	d.theArray = append(d.theArray, value)
	// increment number of values in array by 1
	d.nElems += 1
}

func (d *DArray) display() {
	fmt.Print("[")
	for _, e := range d.theArray {
		fmt.Print(" ", e, " ")
	}
	fmt.Print("]")
}

func (d *DArray) mergeSort() {
	workspace := []int32{}
	d.recMergeSort(workspace, 0, d.nElems-1)
}

func (d *DArray) recMergeSort(workspace []int32, lowerBound int, upperBound int) {
	if lowerBound == upperBound {
		fmt.Println(workspace)
		return
	} else {
		mid := (lowerBound + upperBound) / 2
		d.recMergeSort(workspace, lowerBound, mid)
		d.recMergeSort(workspace, mid+1, upperBound)
		d.merge(workspace, lowerBound, mid+1, upperBound)
	}
}

func (d *DArray) merge(workspace []int32, lowPtr int, highPtr int, upperBound int) {
	j := 0
	lowerBound := lowPtr
	mid := highPtr - 1
	n := upperBound - lowerBound + 1

	// Begin series of while loops...
	for lowPtr <= mid && highPtr <= upperBound {
		if d.theArray[lowPtr] < d.theArray[highPtr] {
			j += 1
			lowPtr += 1
			workspace[j] = d.theArray[lowPtr]
		} else {
			j += 1
			highPtr += 1
			workspace[j] = d.theArray[highPtr]
		}
	}

	for lowPtr <= mid {
		j += 1
		lowPtr += 1
		workspace[j] = d.theArray[lowPtr]
	}

	for highPtr <= upperBound {
		j += 1
		highPtr += 1
		workspace[j] = d.theArray[highPtr]
	}

	for j = 0; j < n; j++ {
		d.theArray[lowerBound+j] = workspace[j]
	}

}

// func main() {
// 	// maxSize := 100

// }
