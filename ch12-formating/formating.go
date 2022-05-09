package main

import "fmt"

type people struct {
	name string
	age  int
}

func main() {
	p := people{"Wilson", 39}
	fmt.Printf("Only type %T\n ", p)
	fmt.Printf("Values and Type %#v \n", p)
	fmt.Printf("Only values  %v", p)
	//Only type main.people
	//Values and Type main.people{name:"Wilson", age:39}
	//Only values  {Wilson 39}%
}
