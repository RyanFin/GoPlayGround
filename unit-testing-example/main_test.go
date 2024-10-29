// math_test.go
package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	// test assertion
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}
