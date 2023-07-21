package main

import (
	"fmt"
)

func main() {
	nums := []int{20, 100, 10, 12, 5, 13}
	//nums := []int{1, 2, 3, 4, 5}
	//nums := []int{2, 1, 5, 0, 4, 6}
	res := increasingTriplet(nums)
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func increasingTriplet(nums []int) bool {
	a := nums[0]
	b := nums[0]

	isExistB := false
	for i := 1; i < len(nums); i++ {
		a = min(nums[i], a)

		if nums[i] > b && nums[i] > a && isExistB == true {
			return true
		}

		if nums[i] > a {
			b = nums[i]
			isExistB = true
		}
	}

	return false
}
