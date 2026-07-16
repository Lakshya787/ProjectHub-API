package main

import "fmt"

func maplit() {
	//var m map[int] string

	m1 := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	v := m1["foo"]
	fmt.Println(v)

}