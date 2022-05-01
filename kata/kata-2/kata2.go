package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Is_valid_ip("127.1.1.aa"))
}

func Is_valid_ip(ip string) bool {
	split := strings.Split(ip, ".")

	if !validSize(split) {
		fmt.Println("validSize", false)
		return false
	}
	if !isNumberEachPos(split) {
		fmt.Println("isNumberEachPos", false)
		return false
	}
	if !isRangeOk(split) {
		fmt.Println("isRangeOk", false)
		return false
	}
	if isLeadingZeros(split) {
		fmt.Println("isLeadingZeros", false)
		return false
	}
	return true
}

func validSize(split []string) bool {
	return len(split) == 4
}

func isNumberEachPos(split []string) bool {
	for _, v := range split {
		if _, err := strconv.Atoi(v); err != nil {
			return false
		}
	}
	return true
}

func isRangeOk(split []string) bool {
	for _, v := range split {
		vint, _ := strconv.Atoi(v)
		if vint > 255 {
			return false
		}
		if vint < 0 {
			return false
		}
	}
	return true
}

func isLeadingZeros(split []string) bool {
	for _, v := range split {
		if strings.HasPrefix(v, "0") {
			if len([]rune(v)) > 1 {
				return true
			}
		}
	}
	return false
}
