package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	resSlice := fizzBuzz(3)

	expected := []string{"1", "2", "Fizz"}

	for i, e := range expected {
		if resSlice[i] != e {
			t.Errorf("want: %v, got: %v", e, resSlice[i])
		}
	}
}
