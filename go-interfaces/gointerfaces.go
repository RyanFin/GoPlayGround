package main

import "fmt"

// Vehicle interface (the contract)
type Vehicle interface {
	Start() string
	Stop() string
	Drive(distance int) string
}

// Car struct (a type that fulfills the contract)
type Car struct {
	make  string
	model string
	year  int
}

// Start method implementation for Car
func (c Car) Start() string {
	return fmt.Sprintf("The %d %s %s has started.", c.year, c.make, c.model)
}

// Stop method implementation for Car
func (c Car) Stop() string {
	return fmt.Sprintf("The %d %s %s has stopped.", c.year, c.make, c.model)
}

// Drive method implementation for Car
func (c Car) Drive(distance int) string {
	return fmt.Sprintf("The %d %s %s has driven %d miles.", c.year, c.make, c.model, distance)
}

func main() {
	myCar := Car{
		make:  "Toyota",
		model: "Corolla",
		year:  2020,
	}

	// Declaring a variable of type Vehicle
	var myVehicle Vehicle = myCar

	// myVehicle can call the methods defined in the Vehicle interface
	fmt.Println(myVehicle.Start())
	fmt.Println(myVehicle.Drive(50))
	fmt.Println(myVehicle.Stop())
}
