package main

import "fmt"

func strlit() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	var arr1 []int
	for i:=0;i<5;i++{
		arr1=append(arr1, i)
	}
	
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))
}