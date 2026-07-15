package main

import "fmt"


func incre(a *[3]int){
	a[0]=100
}

func incrs(a []int ){//need to return a for getting permanent change
	a[0]=100
	a=append(a, 4,5,6)
}



func arrSli() {

	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Printf("A: %v",a)

	b:=[3] int {1,2,3}
	fmt.Printf("B: %v",b)

	var s []int
	s=append(s, 1,2,3)
	fmt.Printf("s: %v",s)

	//s1=make([]int ,3,4)
	//fmt.Printf("s1: %v",s1)
	incre(&a)
	incrs(s)

fmt.Printf("A: %v\n", a)
fmt.Printf("B: %v\n", b)
fmt.Printf("s: %v\n", s)
fmt.Printf("len(A): %d\n", len(a))
fmt.Printf("len(s): %d\n", len(s))
fmt.Printf("A after incre: %v\n", a)
fmt.Printf("s after incrs: %v\n", s)

}