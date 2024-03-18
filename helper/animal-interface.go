package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

// value receiver
func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

//pointer receiver - *
func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	//instantiating value receivers vs instantiating pointer receivers in Go
	animals := []Animal{Dog{}, &Cat{}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
