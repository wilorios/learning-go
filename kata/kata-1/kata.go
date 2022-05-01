package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "1 2 3"
	r := HighAndLow(s)
	fmt.Println(r)
}

func HighAndLow(in string) string {
	var highest, lowest int
	s := strings.Fields(in)
	for i, v := range s {
		ci, _ := strconv.Atoi(v)
		if i == 0 {
			highest = ci
			lowest = ci
		} else {
			if ci > highest {
				highest = ci
			}
			if ci < lowest {
				lowest = ci
			}
		}
	}
	return fmt.Sprint(highest, lowest)
}
