package main

import "fmt"

func operation(x, y int, op func(x, y int) int) int {
	return op(x, y)
}

func add(x, y int) int {
	return x + y
}

func mult(x,y int) int{
	return x*y
}

func multiplier(factor int) func(int) int{
	return func (base int) int{
		return factor*base
	}
}

func highord() {
	//function as input
	fmt.Println("Function as input")
	fmt.Printf("Result: %v",operation(11,67,add))
	fmt.Printf("Result: %v",operation(11,67,mult))

	//function as output
	mult:=multiplier(3)
	fmt.Println(mult(4))
	fmt.Println(mult(12))
}