package main

import "fmt"

func test2(myfunc func(int) int) {
	fmt.Println(myfunc(7))
}

func main() {
	test := func(x int) int {
		return x * 10
	}
	test2(test)
}
