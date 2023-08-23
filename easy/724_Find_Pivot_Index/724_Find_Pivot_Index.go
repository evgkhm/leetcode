package main

import "fmt"

func pivotIndex(nums []int) int {
	sumLeft, sumRight := 0, 0
	for _, val := range nums {
		sumRight += val
	}
	res := -1
	for i := 0; i < len(nums); i++ {
		sumLeft += nums[i]
		if sumLeft == sumRight {
			return i
		}
		sumRight -= nums[i]
	}
	return res
}

func main() {
	var input []int
	input = append(input, 2, 1, -1)
	res := pivotIndex(input)
	fmt.Println(res)
}
