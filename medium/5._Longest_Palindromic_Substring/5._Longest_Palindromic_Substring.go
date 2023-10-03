package main

import "fmt"

func main() {
	s := "ccc"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return string(s[0])
	}

	l, r := 0, 1
	maxLen := 0
	var res string

	for l < len(s)-1 {
		tmpStr := s[l : r+1]
		if isPalindrome(tmpStr) == true {
			if len(tmpStr) > maxLen {
				res = tmpStr
				maxLen = len(tmpStr)
			}
		}

		r++
		if r >= len(s) {
			l++
			r = l + 1
		}
	}

	if res == "" {
		return string(s[0])
	}

	return res
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		left := s[l]
		right := s[r]

		if left != right {
			return false
		}
		l++
		r--
	}

	return true
}
