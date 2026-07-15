package main

import "fmt"

func operator(){
	a:=10
	b:=10
	var c int
	c=a+b//+ ki jagah kuchh bhi daaldo
	fmt.Println("Addition :",c)
	fmt.Println(a*b>a*c)
	fmt.Println(a^b)

	a=3
	a+=4
	fmt.Println(a)
}