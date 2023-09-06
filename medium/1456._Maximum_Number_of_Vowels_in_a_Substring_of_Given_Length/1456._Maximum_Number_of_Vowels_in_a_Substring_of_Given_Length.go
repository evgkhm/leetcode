package main

import (
	"fmt"
)

func main() {
	s := "weallloveyou" //4
	k := 7
	//s := "abciiidef" //3
	//k := 3
	//s := "tnfazcwrryitgacaabwm" //3
	//k := 4
	res := maxVowels(s, k)
	fmt.Println(res)
}

func maxVowels(s string, k int) int {
	l, r := 0, k
	res := 0
	for i := 0; i < r; i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'u' || s[i] == 'o' {
			res++
		}
	}

	maxRes := res

	for r < len(s) {
		if s[l] == 'a' || s[l] == 'e' || s[l] == 'i' || s[l] == 'u' || s[l] == 'o' {
			res--
			if res < 0 {
				res = 0
			}
		}
		if s[r] == 'a' || s[r] == 'e' || s[r] == 'i' || s[r] == 'u' || s[r] == 'o' {
			res++
			if res > maxRes {
				maxRes = res
			}
			if maxRes > k {
				maxRes = k
			}
		}

		l++
		r++
	}

	return maxRes
}
