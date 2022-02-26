//ARRAYS
//ARE COMPORABLE WITH == //Arrays dont share mem
//Arrays dont share mem
package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a, len(a))
	//prints [0 0 0] 3
	a[2] = 30
	fmt.Println(a, len(a))
	//prints [0 0 0] 30
	//you can set a value in a special position
	var b = [10]int{0, 10, 7: 70, 9: 90}
	fmt.Println(b)
	//prints [0 10 0 0 0 0 0 70 0 90]
	arraysDontShareMem()
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
