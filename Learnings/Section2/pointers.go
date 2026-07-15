package main

import "fmt"

func incr(n *int) int{
	return *n+1
}

func pointers() {
	var p *int
	x := 78
	p = &x
	fmt.Printf("Address of x :%d\n",p)
	fmt.Printf("Value of x: %d\n",*p)
	*p=incr(p)
	fmt.Printf("Value of x: %d\n",x)

	pp:=&p
	fmt.Printf("Value of x:%d\n",**pp)
}