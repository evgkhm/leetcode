package main

import "fmt"

func main() {
	//s := "aaaa"
	s := "cbbd" //bb
	//s := "caba" //aba
	//s := "babad" //aba
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return string(s[0])
	}

	maxLen := 0
	var res string

	for i := 0; i < len(s); i++ {
		l, r := i, i // odd
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > maxLen {
				res = s[l : r+1]
				maxLen = r - l + 1
			}
			l--
			r++
		}

		l, r = i, i+1 // even
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > maxLen {
				res = s[l : r+1]
				maxLen = r - l + 1
			}
			l--
			r++
		}
	}

	return res
}
