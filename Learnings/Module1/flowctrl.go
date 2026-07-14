package main

import "fmt"

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

	
	numbers:=[] int {1,2,3}
	for idx,value :=range numbers{
		fmt.Println(idx,value)
	}


}