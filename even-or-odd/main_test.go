// https://medium.com/@rasheed99/how-to-write-table-driven-tests-in-go-8e96ef048cca
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenOrOdd(t *testing.T) {
	// create a  test table for even and odd numbers
	var testTableTests = []struct {
		testName, expectedOutput string
		testInput                int
	}{
		{"test even - 0", "0 - even", 0},
		{"test even - 1", "1 - odd", 1},
		{"test even - 2", "2 - even", 2},
		{"test even - 3", "3 - odd", 3},
		{"test even - 4", "4 - even", 4},
		{"test even - 5", "5 - odd", 5},
		{"test even - 6", "6 - even", 6},
		{"test even - 7", "7 - odd", 7},
		{"test even - 8", "8 - even", 8},
		{"test even - 9", "9 - odd", 9},
		{"test even - 10", "10 - even", 10},
	}

	for _, tt := range testTableTests {
		t.Run(tt.testName, func(t *testing.T) {
			result := evenOrOdd(tt.testInput)

			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}
