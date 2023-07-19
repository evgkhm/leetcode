package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 121
	res := isPalindrome(x)
	fmt.Println(res)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	l, r := 0, len(str)-1
	middle := (len(str) - 1) / 2
	for l <= middle {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}
