package main

import (
	"fmt"
	"goproject/Learnings/module1/packages/muth"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("my first word")
	fmt.Printf("addition using math package: %v\n", muth.Add(1, 2))
	fmt.Println(uuid.New().String())
}