package main

import (
	"fmt"
	"math"
)

func main() {
	mathConvertion(1.4)
}

func mathConvertion(intput float64) {
	//fmt.Println(float64(2))
	fmt.Println(math.Floor(intput))
	fmt.Println(math.Ceil(intput))
	fmt.Println(math.Round(intput))
	fmt.Println(math.RoundToEven(intput))
}
