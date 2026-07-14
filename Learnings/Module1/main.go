package main

import "fmt"

/*import (
	"fmt"
	"goproject/Learnings/module1/packages/muth"
	"github.com/google/uuid"
)*/

func main() {
	//operator()
	Hello(2,3)
	Som(2,3,41)
	greet:=func (name string){
		fmt.Println("Hello , from GFG to ",name)
	}
	greet("Lakshya")
}