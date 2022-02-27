package main

import (
	"errors"
	"fmt"
)

//WAY #1 TO CALL ERROR "errors.New"
func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("Error #1 only even numbers are processed")
	}
	return i * 2, nil
}

//WAY #2 TO CALL ERROR "fmt.Errorf"
func doubleEven2(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("Error #2 %d isn't an even number", i)
	}
	return i * 2, nil
}
func main() {
	a := 11
	remainder, err := doubleEven(a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("test 1 remainder ok ", remainder)
	}
	remainder, err = doubleEven2(a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("test 2 remainder ok ", remainder)
	}
}
