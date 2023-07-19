package main

import "fmt"

func main() {
	//TEST1
	nums := make([]int, 6)
	nums[0] = 2
	nums[1] = 3
	nums[2] = 1
	nums[3] = 2
	nums[4] = 4
	nums[5] = 3
	target := 7

	//TEST2
	//nums := make([]int, 3)
	//nums[0] = 1
	//nums[1] = 4
	//nums[2] = 4
	//target := 4

	//TEST3
	//nums := make([]int, 8)
	//nums[0] = 1
	//nums[1] = 1
	//nums[2] = 1
	//nums[3] = 1
	//nums[4] = 1
	//nums[5] = 1
	//nums[6] = 1
	//nums[7] = 1
	//target := 11

	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSubArrayLen(target int, nums []int) int {
	currIdx := 0
	minLenSubarray := len(nums)
	lenSubArray := 0

	sum, counter := 0, 0
	deleteIdx := 0
	for currIdx < len(nums) {
		sum += nums[currIdx]
		counter++

		for sum >= target {
			lenSubArray = counter
			minLenSubarray = min(lenSubArray, minLenSubarray)

			sum -= nums[deleteIdx]
			counter--
			deleteIdx++
		}

		currIdx++
	}
	return min(lenSubArray, minLenSubarray)
}
