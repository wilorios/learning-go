package main

import "fmt"

func main() {
	var a int = 10

	var b float64 = 30.2
	//go doesnt have explicit conversion
	c := float64(a) + b
	var d int = a + int(b)
	fmt.Println(a, b, c, d)
	//prints
}
