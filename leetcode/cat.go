package main

import "fmt"

type Cat struct {
	IsSleeping bool
}

func (c Cat) Sleep() {
	c.IsSleeping = true
}

func main() {
	b := Cat{}
	b.Sleep()
	fmt.Println(b.IsSleeping)
}
