package main

import "fmt"

type shape interface {
	getArea() float64 // both square and triangle fulfil the shape interface
}

type triangle struct {
	base, height float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	fmt.Println("hello world")
	square := square{sideLength: 3}
	triangle := triangle{base: 23, height: 70}

	printArea(square)
	printArea(triangle)
}
