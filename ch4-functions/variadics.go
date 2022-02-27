package main

import "fmt"

func main() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	//ojo el ... es requerido de lo contrario saca error cannot use a (type []int) as type int in //argument to addTo
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}
