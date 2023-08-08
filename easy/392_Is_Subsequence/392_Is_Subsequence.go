package main

import "fmt"

func main() {
	s := "acb"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	idx := 0

	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(t); i++ {
		if t[i] == s[idx] {
			idx++

			if idx == len(s) {
				return true
			}
		}
	}

	return false
}
