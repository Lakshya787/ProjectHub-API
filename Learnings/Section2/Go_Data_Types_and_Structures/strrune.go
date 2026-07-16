package main

import "fmt"

func strrune(){
	fmt.Println("String and Rune")
	s:="Hello World 🗿"
	r:=[]rune (s)
	fmt.Println(len(s))
	fmt.Println(len(r))

	for idx,char:= range s{
		fmt.Printf("String character: %c, and index: %d with ASCII value: %v\n",char,idx,char)
	}
	
	for idx,char:= range r{
		fmt.Printf("Character: %c | Byte Index: %d | Unicode: U+%04X\n", char, idx, char)
	}
}