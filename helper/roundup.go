package main

import (
	"fmt"
	"math"
)

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func RoundPlus(f float64, places int) (float64) {
	shift := math.Pow(10, float64(places))
	return Round(f * shift) / shift;
}

func main() {
	fmt.Println(Round(123.4999))
	fmt.Println(Round(123.5))
	fmt.Println(Round(123.999))

	fmt.Println(RoundPlus(123.554999, 3))
	fmt.Println(RoundPlus(123.555555, 3))
	fmt.Println(RoundPlus(123.558, 2))
}
