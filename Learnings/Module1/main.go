package main

import "fmt"

/*import (
	"fmt"
	"goproject/Learnings/module1/packages/muth"
	"github.com/google/uuid"
)*/

func main() {
	//operator()
	/*Hello(2,3)
	Som(2,3,41)
	greet:=func (name string){
		fmt.Println("Hello , from GFG to ",name)
	}
	greet("Lakshya")*/
	fmt.Println("Code")
	//loops()
	/*defer demo(1)
	demo(2)
	defer demo(3)*/
	//demoGOTO()
	fmt.Println(globalProtanogonist)
	var sen="I love Sentry"
	fmt.Println(sen)//local

	if (true){
		var gg="Showdown"
		fmt.Println(gg)//block
	}
		counter := closuree()

	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	anotherCounter := closuree()

	fmt.Println(anotherCounter()) // 1
	fmt.Println(counter())        // 4
}