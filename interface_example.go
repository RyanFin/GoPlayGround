package main

import "fmt"

type InterfaceExample interface {
	FirstFunction(a int, b int) int
	SecondFunction(a string, b int) int
}

func main() {
	fmt.Println(FirstFunction(4, 5))
}

func FirstFunction(a int, b int) int {
	return a + b
}
