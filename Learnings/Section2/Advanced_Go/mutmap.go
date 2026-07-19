package main

import "fmt"

func change(m map[string] int){
	m["apple"]=50
}


func mutmap() {
	m:=map[string] int {
		"Apple":1,
		"Orange":2,
	}

	fmt.Println(m)
	//print an element
	fmt.Println(m["apple"])
	m["apple"]=10

	if val,ok:=	m["apple"];ok{
		fmt.Println("old value",val)
		m["apple"]=25
	}
	fmt.Println(m["apple"])
	delete(m,"apple")
	fmt.Println(m["apple"])
	change(m)
	fmt.Println(m["apple"])
	
}