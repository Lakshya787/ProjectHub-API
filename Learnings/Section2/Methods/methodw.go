package main

import "fmt"

type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) Area() int {
	return r.Width * r.Length
}

func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Length *= factor
}

func methodw() {
	rectangle := Rectangle{
		Length: 3,
		Width:  5,
	}

	fmt.Println(rectangle.Area())

	rectangle.Scale(2)
	fmt.Println(rectangle.Area())
}