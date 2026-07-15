package main

import "fmt"

func maps() {
	var m map[int]string

	if m == nil {
		fmt.Println("Map is nil")
	}
	ml:=make(map[int]string)
	ml[100]="Udemy"
	val:=ml[100]
	fmt.Printf("Map Value: %v\n",val)

	val,exists:=m[200]
	if !exists{
		fmt.Printf("doesnt exist \n")
	}else{
		fmt.Printf("Map value: %v\n",val)
	}

}