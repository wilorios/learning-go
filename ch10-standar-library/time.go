package main

import (
	"fmt"
	"time"
)

func main() {
	d := 2*time.Hour + 30*time.Minute
	fmt.Println("time duration ", d)
	//prints time duration 2h30m0s
	t, err := time.Parse("2006-02-01 15:04:05 -0700", "2016-13-03 00:00:00 +0000")
	if err != nil {
		fmt.Println("error parsing time ")
	} else {
		fmt.Println("parsing time ", t)
		//prints parsing time 2016-03-13 00:00:00 +0000 +0000
	}
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
	//prints March 13, 2016 at 12:00:00AM UTC
}
