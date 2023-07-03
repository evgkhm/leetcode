package main

import (
	"fmt"
)

func main() {
	s := "ceabaacb"
	//s := "aaabbbcc"
	//s := "bbcebab" //2
	res := minDeletions(s)
	fmt.Println(res)
}

func minDeletions(s string) int {
	m := make(map[string]int)
	counter := 0
	for _, val := range s {
		_, ok := m[string(val)]
		if !ok {
			m[string(val)] = 1
		} else {
			counter = m[string(val)]
			counter++
			m[string(val)] = counter
		}
	}
	var slice []int
	for _, val := range m {
		slice = append(slice, val)
	}
	res := 0
	removeDuplicate(slice, &res)

	return res
}

func removeDuplicate(slice []int, res *int) {
	m := make(map[int]struct{})

	for i := 0; i < len(slice); i++ {
		_, ok := m[slice[i]]
		if !ok {
			m[slice[i]] = struct{}{}
		} else {
			slice[i]--
			if slice[i] < 0 {
				continue
			}

			*res++
			removeDuplicate(slice, res)
		}
	}
}
