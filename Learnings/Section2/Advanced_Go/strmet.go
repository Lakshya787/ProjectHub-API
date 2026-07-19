package main

import "fmt"

type Car struct {
	Name  string
	Type  string
	Brand string
	Years int
}

//Method Years
func (c Car) operational(){
	if c.Years>15{
		fmt.Printf("Car with name %s and Brand %s is not operational",c.Name,c.Brand)
	}else{
				fmt.Printf("Car with name %s and Brand %s and %v years of operation is left",c.Name,c.Brand,15-c.Years)
	}
}

func strmet() {
	c := Car{"Model S", "EV", "Tesla", 3}
	fmt.Printf("Name of the car: %s\n",c.Name)
	c.operational()
}