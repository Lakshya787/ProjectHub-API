package main

import "fmt"

func constants() {
	const a = 10
	const b float64 = 3.14
	fmt.Printf("a=%v,Type a=%T\nb=%v,Type b=%T\n",a,a,b,b)
	const (
		L = 5
		W = 10
		A = L * W
	)
	fmt.Printf("Area=%v, Type Area=%T",A,A)

	const (
		One = iota
		Two
		Three
	)
	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Three)

	z:=float64(a)
	fmt.Printf("Type z: %T",z)
}