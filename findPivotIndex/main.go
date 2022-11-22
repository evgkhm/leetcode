package main

import "fmt"

func pivotIndex(nums []int) int {
	res := -1
	for pivot := 0; pivot < len(nums); pivot++ {
		left := 0
		right := 0

		for j := 0; j < pivot; j++ {
			left += nums[j]
		}

		for j := len(nums) - 1; j > pivot; j-- {
			right += nums[j]
		}

		if left == right {
			res = pivot
			break
		}
	}
	return res
}

func main() {
	var input []int
	input = append(input, -1, -1, 0, 1, 1, 0)
	res := pivotIndex(input)
	fmt.Println(res)
}
