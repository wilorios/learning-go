package main

import "fmt"

//EXAMPLE 1 //interface "forma"
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rect struct {
	width  float64
	height float64
	color  string
}

//implementation of interface
func (r Rect) Area() float64 {
	return r.width * r.height
}

//implementation of interface
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

//EXAMPLE 2
//empty interface (zero methods)
//its like "Println" we can receive any value
func emptyInterface2(i interface{}) {
	//type assertion
	switch i.(type) {
	case string:
		fmt.Println("string ", i.(string))
	case int:
		fmt.Println("int", i)
	default:
		fmt.Println("default", i)
	}
}

//EXAMPLE 3 //Embedding interfaces
type Color interface {
	Exacode() string
}

//implementation of interface
func (r Rect) Exacode() string {
	return r.color
}

type EmbeddingInt interface {
	Shape
	Color
}

func main() {
	//************EXAMPLE 1
	//we can create a variable of interface
	var s Shape
	//this variable could be assigned to the //implementation of the interface
	s = Rect{5.0, 4.0, "red"}
	// the previous variable can be Shape or Rect //thats it polymorphism
	r := Rect{5.0, 4.0, "red"}
	fmt.Printf("type of s is %T\n", s) //prints: type of s is main.Rect fmt.Printf("value of s is %v\n", s) //prints: value of s is {5 4 red}

	fmt.Println("Area ", s.Area()) //prints: Area 20
	fmt.Println("s==r", s == r)
	//prints: s==r true
	//prints true because has the same struct of
	//type Rect and the same dynamic value 5.0 and 4.0 and red
	//*************EXAMPLE 2
	emptyInterface2(20)
	//prints: int 20
	emptyInterface2(s)
	//prints: default {5 4 red}
	//************EXAMPLE 3
	var embedding EmbeddingInt = r
	fmt.Printf("type of embedding is %T\n", embedding)
	//prints: type of embedding is main.Rect
}
