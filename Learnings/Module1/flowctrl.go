package main

import (
	"fmt"
	"strconv"
)

func loops() {
	fmt.Println("Basic")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", i)
	}

	fmt.Println("While")
	x := 5
	for x < 10 {
		fmt.Println(x)
		x++
	}

	//infinite loop
	fmt.Println("infinite loop...")
	x=0
	for{
		if x>10{
			break
		}
		fmt.Println(x)
		x++
	}

	fmt.Println("Index and Value")
	numbers:=[] int {1,2,3}
	for idx,value :=range numbers{
		fmt.Println(idx,value)
	}

	y:=5
	if y>5{
		fmt.Println("y is greater than 5")
	}else if y==5{
		fmt.Println("y is equal to 5")
	}else{
		fmt.Println("y is less than equal to 5")
	}

	fmt.Println("Short hand ")
	if _,e :=strconv.Atoi("345");e!=nil{
		fmt.Println(e)
	}

	day:=2
	switch day{
	case 1:
		fmt.Println("It's a monday")

	case 2:
		fmt.Println("It's Tuesday")
	default:
		fmt.Println("Other day")
	}

	switch day{
	case 1,2,3:
		fmt.Println("Its a week start")
	case 4,5,6:
		fmt.Println("2nd half")
	default:
		fmt.Println("end")
	}

	var p interface{}=6

	switch p.(type){
	case int:
		fmt.Println("X is integer")
	case string:
		fmt.Println("X is a string")
	default:
		fmt.Println("Idk man wth is wrong here")
	}


}