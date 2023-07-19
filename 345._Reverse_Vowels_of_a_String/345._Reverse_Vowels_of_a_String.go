package main

import (
	"fmt"
)

func main() {
	s := "leetcode"
	res := reverseVowels(s)
	fmt.Println(res)
}

func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	foundL, foundR := false, false

	res := []rune(s)
	for l <= r {
		if res[l] == 'a' || res[l] == 'A' || res[l] == 'e' || res[l] == 'E' || res[l] == 'i' ||
			res[l] == 'I' || res[l] == 'o' || res[l] == 'O' || res[l] == 'u' || res[l] == 'U' {
			foundL = true
		} else {
			l++
		}

		if res[r] == 'a' || res[r] == 'A' || res[r] == 'e' || res[r] == 'E' || res[r] == 'i' ||
			res[r] == 'I' || res[r] == 'o' || res[r] == 'O' || res[r] == 'u' || res[r] == 'U' {
			foundR = true
		} else {
			r--
		}

		if foundL == true && foundR == true {
			res[l], res[r] = res[r], res[l]
			foundL, foundR = false, false
			l++
			r--
		}
	}
	return string(res)
}
