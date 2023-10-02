package main

import "fmt"

func main() {
	s := "dvdf" // 3
	//s := "abcabcbb" // 3
	//s := "bbbbb" // 1
	//s := "pwwkew" // 3
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	l, r := 0, 0
	uniqSymbols := make(map[uint8]int)
	res := 0

	for i := 0; i < len(s); i++ {
		uniqSymbols[s[i]] += 1

		if uniqSymbols[s[i]] > 1 {
			val, _ := uniqSymbols[s[i]]
			for val > 1 {
				uniqSymbols[s[l]] -= 1
				val = uniqSymbols[s[i]]
				l += 1
			}
		}

		res = max(r-l+1, res)
		r++
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
