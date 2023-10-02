package main

import "fmt"

func main() {
	nums := []int{1, 0, 1, -4, -3} // false
	//nums := []int{1, 2, 3, 4} // false
	//nums := []int{3, 1, 4, 2} // true
	//nums := []int{3, 5, 0, 3, 4} // true
	//nums := []int{-2, 1, 1, -2, 1, 1} // false
	//nums := []int{3, 1, 4, 2}
	fmt.Println(find132pattern(nums))
}

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	minSlice := make([]int, len(nums))
	minSlice[0] = nums[0]
	for i := 1; i < len(minSlice); i++ {
		minSlice[i] = min(minSlice[i-1], nums[i])
	}

	var stack []int
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > minSlice[i] {
			for len(stack) > 0 && stack[len(stack)-1] <= minSlice[i] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 && stack[len(stack)-1] < nums[i] {
				return true
			}
			stack = append(stack, nums[i])
		}
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
