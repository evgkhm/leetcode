package main

import (
	"fmt"
)

func main() {
	//arr := [...]int{1, 2, 3, 4}
	//difference := 1
	arr := [...]int{1, 5, 7, 8, 5, 3, 4, 2, 1}
	difference := -2
	res := longestSubsequence(arr, difference)
	fmt.Println(res)
}

func longestSubsequence(arr [9]int, difference int) int {
	m := make(map[int]int)
	res := 0
	for _, val := range arr {
		a := val

		beforeA := a - difference

		_, ok := m[beforeA]
		if ok {
			m[a] = m[beforeA] + 1
		} else {
			m[a] = 1
		}

		res = max(res, m[a])
	}

	return res
}

// MINIMALISTIC SOLUTION
//func longestSubsequence(arr [9]int, difference int) int {
//	m := make(map[int]int, len(arr))
//	res := 0
//	for _, val := range arr {
//		m[val] = m[val-difference] + 1
//		res = max(res, m[val])
//	}
//	return res
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
