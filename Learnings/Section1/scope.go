package main 

import "fmt"

var globalProtanogonist="Lakshya"

func scope(){
	fmt.Println("Scope")
}

func closuree() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}