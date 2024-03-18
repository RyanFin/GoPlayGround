package main

import (
	"fmt"
)

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	names := []string{"stanley", "david", "oscar"}

	// convert []interface array to []string array
	vals := make([]interface{}, len(names))

	for i, e := range names {
		vals[i] = e
	}

	PrintAll(vals)
}
