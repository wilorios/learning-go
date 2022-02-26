package main

import "fmt"

func main() {
	x := 3
	if x > 1 {
		fmt.Println("Nice")
	}
	//complete style statement
	for i := 0; i < x; i++ {
		fmt.Println(i)
	}
	//conditional only statement
	for x < 4 {
		fmt.Println("conditional->", x)
		x++
	}
	//infinite style
	//for {
	//}

	//for-range
	evenVals := []int{2, 4, 6, 8}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	uniqueNames := map[string]bool{"Fred": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
	for k, v := range uniqueNames {
		fmt.Println(k, v)
	}

	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

}
