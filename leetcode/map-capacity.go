package main

import "fmt"

func main() {
	// create a map first without capacity
	animalMap := make(map[int]string)

	animalMap[0] = "cat"
	animalMap[1] = "dog"
	animalMap[2] = "bat"

	// output entire map
	fmt.Println(animalMap)

	// access only one element in the map
	fmt.Println(animalMap[0])

	// create a map with capacity
	smashMap := make(map[int]string)

	smashMap[0] = "Yoshi"
	smashMap[1] = "Mario"
	smashMap[2] = "Bowser"
	smashMap[3] = "Peach"
	smashMap[4] = "Daisy"

	fmt.Println(smashMap)

}
