package main

import "fmt"

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}

func main() {
	var input []int
	input = append(input, 1, 2, 3, 4)
	res := runningSum(input)
	fmt.Println(res)
}
