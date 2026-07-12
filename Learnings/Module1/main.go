package main

import (
	"fmt"
	"goproject/Learnings/module1/packages/muth"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("my first word")
	fmt.Printf("addition using math package: %v\n", muth.Add(1, 2))
	fmt.Println(uuid.New().String())
	var age int //
	var fl float32
	var b bool
	var str string
	var arr[3] int
	var slice [] int
	var m map[string] int
	var ptr *int
	fmt.Println("int: %d",age)
	fmt.Println("float: %f",fl)	
	fmt.Println("boolean: %v",b)	
	fmt.Println("int: %s",str)
	fmt.Println("int: %v",arr)
	fmt.Println("int: %v",slice)
	fmt.Println("Pointer: %v",m)
	fmt.Println("Pointer: %v",ptr)

}