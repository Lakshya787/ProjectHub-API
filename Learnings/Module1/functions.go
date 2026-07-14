package main

import (
	"fmt"
	"strconv"
)

func Hello(a int , b int )int {//can also write like a,b int
	fmt.Printf("Hello, GfG people! The sum of two variables is: %v\n",a+b)
	return a+b
}

func Som(numbers ...int) int {
	total:=0
	for i:= range numbers{
		total+=numbers[i]
	}
	return total
}

func Swap(x,y int) (a int,b int){
	a=y
	b=x
	return 
}

func Area(a,b float64) (area float64,p float64){
	area=a*b
	p=2*(a+b)
	return
}

func Counter(a int) func() int{
	count:=10
	return func() int{
		count+=a
		return count
	}
}

func Convert( str string) (int,error){
	return strconv.Atoi(str)
}