package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = 30.2
	//go doesnt have explicit conversion
	c := float64(a) + b
	var d int = a + int(b)
	e := (a != d)

	const f = 10 * 20

	fmt.Println(a, b, c, d, e, f)
	//prints 10 30.2 40.2 40 true 200
}
