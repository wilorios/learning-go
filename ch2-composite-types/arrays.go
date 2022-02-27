//ARRAYS
//ARE COMPORABLE WITH ==
//Arrays dont share mem
//Arrays the function change the data
//hence we dont need to use the pointer reference
package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a, len(a))
	//prints [0 0 0] 3
	a[2] = 30
	fmt.Println(a, len(a))
	//prints [0 0 30] 3
	//you can set a value in a special position
	var b = [10]int{0, 10, 7: 70, 9: 90}
	fmt.Println(b)
	//prints [0 10 0 0 0 0 0 70 0 90]
	arraysDontShareMem()
	arraysDontShareMem2(a)
	fmt.Println(a, b)
	//prints [0 0 30] [0 10 0 0 0 0 0 70 0 90]
}

func arraysDontShareMem() {
	var a = [2]int{1, 2}
	var b = a
	fmt.Println(b, a == b)
	//[1 2] true
	b[0] = 3
	fmt.Println(a, b, a == b)
	//[1 2] [3 2] false
}

func arraysDontShareMem2(a [3]int) {
	var b = a
	fmt.Println(b, a == b)
	//[0 0 30] true
	b[0] = 3
	fmt.Println(a, b, a == b)
	//[0 0 30] [3 0 30] false
}
