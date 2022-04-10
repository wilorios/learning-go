package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println("1->", x, y, z)

	y = append(y, 30, 40, 50)
	fmt.Println("2->", x, y, z)

	x = append(x, 60)
	fmt.Println("3->", x, y, z)
	z = append(z, 70)
	fmt.Println("4->", x, y, z)

	for i := 0; i < 10; i++ {
		fmt.Println(i % 3)
	}
}
