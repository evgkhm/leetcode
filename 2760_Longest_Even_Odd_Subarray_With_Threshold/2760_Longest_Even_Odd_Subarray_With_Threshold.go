package main

import "fmt"

func main() {
	nums := make([]int, 2)
	nums[0] = 8
	nums[1] = 4

	threshold := 6

	res := longestAlternatingSubarray(nums, threshold)
	fmt.Println(res)
}

func longestAlternatingSubarray(nums []int, threshold int) int {
	currSize, maxSize, prevVal, start := 0, 0, 1, false
	for _, val := range nums {
		if start == false && val%2 == 0 && val <= threshold {
			start = true
		}

		if start == true && prevVal%2 != val%2 && val <= threshold {
			currSize++
			prevVal = val % 2
			if currSize > maxSize {
				maxSize = currSize
			}
			continue
		} else if val%2 == 0 && val <= threshold {
			currSize = 1
			prevVal = val % 2
			if currSize > maxSize {
				maxSize = currSize
			}
			continue
		}
		currSize, prevVal, start = 0, val%2, false
	}

	return maxSize
}
