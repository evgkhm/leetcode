package main

import (
	"fmt"
)

func main() {
	//s := "cbaebabacd"
	//p := "abc"

	s := "abab"
	p := "ab"

	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	symbolsMap := make(map[rune]int)

	for _, val := range p {
		symbolsMap[val] += 1
	}

	var res []int
	l := 0

	for r := 0; r < len(s); r++ {
		symbolsMap[rune(s[r])] -= 1

		if symbolsMap[rune(s[r])] == 0 {
			delete(symbolsMap, rune(s[r]))
		}

		if len(symbolsMap) == 0 {
			res = append(res, l)
		}

		if r >= len(p)-1 {
			symbolsMap[rune(s[l])] += 1
			if symbolsMap[rune(s[l])] == 0 {
				delete(symbolsMap, rune(s[l]))
			}

			l += 1 // move left pointer
		}
	}

	return res
}
