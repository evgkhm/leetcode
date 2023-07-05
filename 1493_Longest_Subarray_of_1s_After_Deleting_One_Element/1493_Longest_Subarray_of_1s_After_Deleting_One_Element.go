package main

import "fmt"

func main() {
	//TEST1
	//nums := make([]int, 4)
	//nums[0] = 1
	//nums[1] = 1
	//nums[2] = 0
	//nums[3] = 1

	//TEST2
	//nums := make([]int, 9)
	//nums[0] = 0
	//nums[1] = 1
	//nums[2] = 1
	//nums[3] = 1
	//nums[4] = 0
	//nums[5] = 1
	//nums[6] = 1
	//nums[7] = 0
	//nums[8] = 1

	//TEST3
	//nums := make([]int, 3)
	//nums[0] = 1
	//nums[1] = 1
	//nums[2] = 1

	//TEST56
	//nums := make([]int, 4)
	//nums[0] = 0
	//nums[1] = 0
	//nums[2] = 1
	//nums[3] = 1

	nums := make([]int, 14)
	nums[0] = 1
	nums[1] = 0
	nums[2] = 1
	nums[3] = 1
	nums[4] = 1
	nums[5] = 1
	nums[6] = 1
	nums[7] = 1
	nums[8] = 0
	nums[9] = 1
	nums[10] = 1
	nums[11] = 1
	nums[12] = 1
	nums[13] = 1

	res := longestSubarray(nums)
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubarray(nums []int) int {
	zeroCounter := 0
	longest := 0
	right, left := 0, 0

	for right < len(nums) {
		longest = max(longest, right-left)

		if nums[right] == 0 {
			zeroCounter++
		}

		for zeroCounter > 1 {
			if nums[left] == 0 {
				zeroCounter--
			}
			left++
		}

		right++
	}
	return max(longest, right-left) - 1
}
