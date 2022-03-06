package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p := person{
		firstName: "Benji",
		lastName:  "Rios",
		age:       8}
	output := p.string()
	fmt.Println(output)
}

//the difference with the functions is "the receiver"
//in that case the receiver is (p person)
func (p person) string() string {
	return fmt.Sprintf("%s %s, age %d", p.firstName, p.lastName, p.age)
}
