package main

import "fmt"

func main() {
	var mySavings float32 = 3000

	interestToBeAdded := InterestRateCalculator(mySavings)
	fmt.Println(interestToBeAdded)
}

// return the interst amount to be applied to a savings account
func InterestRateCalculator(rate float32) float32 {
	fmt.Println("start")
	var interest float32
	var ret float32
	if rate >= 0 {
		if rate > 0 && rate <= 500 {
			interest = 1
			ret = (interest / 100) * rate
		}

		if rate > 500 && rate <= 1000 {
			interest = 2
			ret = (interest / 100) * rate
			fmt.Println("ret: ", ret)
		}

		if rate > 1000 {
			interest = 3
			ret = (interest / 100) * rate
		}
	}

	return ret
}
