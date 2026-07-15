package main

import "fmt"

func demoGOTO(){
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			 if i==1 && j==1{
				goto end
			 }
			 fmt.Println(i,j)
		}
	}
	end:
	fmt.Println("Exit of loop")
}
