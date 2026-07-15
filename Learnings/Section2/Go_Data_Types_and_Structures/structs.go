package main

import "fmt"

func structs() {
	fmt.Println("Structs")

	type Car struct {
		Name  string
		Type  string
		Brand string
		Years int
	}

	var c Car

	c = Car{Name: "Model S", Type: "EV", Brand: "Tesla", Years: 3}
	fmt.Println(c.Brand)

	c = Car{"Model X", "EV", "Tesla", 1}
	fmt.Println(c.Brand)

	c = Car{Type: "EV"}
	fmt.Println(c)

	type Power struct {
		HP   int
		Type string
	}

	type CAR struct {
		Name  string
		Type  string
		Brand string
		Years int
		Power // Embedded struct
	}

	var c2 CAR

	c2 = CAR{
		"Model XX",
		"EV",
		"Tesla",
		2,
		Power{677, "Electric"},
	}

	fmt.Println(c2.Brand)
	fmt.Println(c2.HP)         // Promoted field
	fmt.Println(c2.Power.Type) // Type inside embedded Power
	p:=struct{
		Name string
		Age int
	}{
		Name:"GfG",
		Age:3,
	}
	fmt.Println(p.Name)
}