package pkg

import (
	"fmt"
	"testing"
)

func TestInterestCalculate(t *testing.T) {
	var mySavingsBalance float32 = 100

	res := CalculateInterest(mySavingsBalance)
	expected := 1
	if res != float32(expected) {
		t.Errorf("want %v , got %v", expected, res)
	}
}

func TestInterestCalculateV2(t *testing.T) {
	var mySavingsBalance float32 = 1000

	res := CalculateInterest(mySavingsBalance)
	expected := 20
	if res != float32(expected) {
		t.Errorf("want %v , got %v", expected, res)
	}
}

func TestInterestCalculateV3(t *testing.T) {
	var mySavingsBalance float32 = 3000

	res := CalculateInterest(mySavingsBalance)
	expected := 60
	if res != float32(expected) {
		t.Errorf("want %v , got %v", expected, res)
	}
}

func TestNegativeInput(t *testing.T) {
	var mySavingsBalance float32 = -300
	res := CalculateInterest(mySavingsBalance)
	fmt.Println(res)
}
