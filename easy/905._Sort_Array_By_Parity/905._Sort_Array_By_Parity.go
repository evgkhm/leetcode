package main

import "fmt"

func main() {
	nums := []int{0, 1, 3}
	fmt.Println(sortArrayByParity(nums))
}

func sortArrayByParity(nums []int) []int {
	res := make([]int, len(nums))
	l, r := 0, len(nums)-1

	for _, val := range nums {
		if val%2 == 0 { //четные
			res[l] = val
			l++
		} else {
			res[r] = val
			r--
		}
	}

	return res
}
