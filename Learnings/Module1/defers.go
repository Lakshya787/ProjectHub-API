package main

import (
	"fmt"
	"os"
)

func demo(i int) {
	fmt.Println("Hello Lakshya",i)
}

func capture(){
	for i:=0;i<3;i++{
		defer fmt.Println(i)
	}
}

func OpenCheck(){
	file,err:=os.Open("temp.txt")
	if err!=nil{
		fmt.Errorf("error opening temp.txt: %v",err)
	}
	defer file.Close()
}

