package main

import (
	"fmt"
	"goproject/Learnings/module1/packages/muth"

	"github.com/google/uuid"
)

func variable() {
	fmt.Println("my first word")
	fmt.Printf("addition using math package: %v\n", muth.Add(1, 2))
	fmt.Println(uuid.New().String())

	var age int
	var fl float32
	var b bool
	var str string
	var arr [3]int
	var slice []int
	var m map[string]int
	var ptr *int

	fmt.Printf("int: %d\n", age)
	fmt.Printf("float: %f\n", fl)
	fmt.Printf("boolean: %v\n", b)
	fmt.Printf("string: %q\n", str)
	fmt.Printf("array: %v\n", arr)
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("map: %v\n", m)
	fmt.Printf("pointer: %v\n", ptr)

	//shortHand initialization
	a :=3
	fmt.Printf("a= %d\n",a)
	arr1 := [3]int{1,2,3}
	fmt.Printf("arr =%v\n",arr1)
}