package main

import "fmt"

func fizzBuzz(n int) []string {

	var ret []string

	for i := 1; i <= n; i++ {
		fmt.Println(i % 15)

		if i%3 == 0 && i%5 == 0 {
			ret = append(ret, "FizzBuzz")
		} else if i%3 == 0 {
			ret = append(ret, "Fizz")
		} else if i%5 == 0 {
			ret = append(ret, "Buzz")
		} else {
			ret = append(ret, fmt.Sprint(i))
		}
	}

	return ret

}

func main() {
	fmt.Println(fizzBuzz(30))
}
