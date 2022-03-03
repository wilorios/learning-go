package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	p := person{
		name: "Wilson",
		age:  39,
		pet:  "dog",
	}
	i := 1
	fmt.Println("1->", p, i)
	//print 1-> {Wilson 39 dog} 1
	modifyFails(p, i)
	fmt.Println("2->", p, i)
	//print 2-> {Wilson 39 dog} 1
	modifyOk(&p, &i)
	fmt.Println("3->", p, i)
	//print 3-> {Benji 39 dog} 2
}

func modifyFails(p person, i int) {
	p.name = "Benji"
	i = 2
}

func modifyOk(p *person, i *int) {
	p.name = "Benji"
	*i = 2
}
