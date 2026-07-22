package main

import (
	"fmt"
)

//for using interface you have to implement every method
type Car interface {
	StartEngine() string
	Drive() string
}
type Sedan struct {
	Brand  string
	Length int
}

type SUV struct {
	Brand  string
	Height int //test
}

func (s Sedan) StartEngine() string {
	return fmt.Sprintf("%s Brand car has started engine silently\n",s.Brand)
}
func (s Sedan) Drive() string{
	return fmt.Sprintf("%s Brand sedan has length %v can be driven smoothly on highways\n",s.Brand,s.Length)
}

func (s SUV) StartEngine() string {
	return fmt.Sprintf("%s Brand car has started engine silently\n",s.Brand)
}
func (s SUV) Drive() string{
	return fmt.Sprintf("%s Brand sedan has length %v can be driven smoothly on highways\n",s.Brand,s.Height)
}

func DefineCar(c Car){
	switch tp:=c.(type){
	case Sedan:
		fmt.Println(tp.Drive())
	case SUV:
		fmt.Println(tp.StartEngine())
	default:
		fmt.Println("type didnt match")
	
	}
}

func interfacestyas() {
	/*var car Car
	car=Sedan{Brand: "Honda",Length: 15}
	DefineCar(car)
	var output interface{}
	output = 45
	i, ok := output.(int)
	if !ok {
		fmt.Println("Unable to convert output")
	}
	fmt.Printf("Integer %v\n",i)
	output="GFG"
	str,ok:=output.(int)
	if !ok{
		fmt.Println("Unable to convert output")
	}
	fmt.Printf("string %d\n",str)*/
	var c Car
	c=Sedan{Brand: "honda",Length: 15}
	DefineCar(c)
	c=SUV{Brand: "JEEP",Height: 250}
	DefineCar(c)
}
