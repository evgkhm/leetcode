package main

import (
	"fmt"
)

func main() {
	s := "rat"
	t := "car"
	res := isAnagram(s, t)
	fmt.Println(res)
}

func isAnagram(s string, t string) bool {
	var s1 [26]int
	var s2 [26]int

	for _, val := range s {
		idx := val - 'a'
		s1[idx]++
	}

	for _, val := range t {
		idx := val - 'a'
		s2[idx]++
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
