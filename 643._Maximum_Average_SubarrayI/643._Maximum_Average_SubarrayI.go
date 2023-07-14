package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 6)
	nums[0] = 1
	nums[1] = 12
	nums[2] = -5
	nums[3] = -6
	nums[4] = 50
	nums[5] = 3
	k := 4

	//nums := make([]int, 1)
	//nums[0] = -1

	//nums := [...]int{8860, -853, 6534, 4477, -4589, 8646, -6155, -5577, -1656, -5779, -2619, -8604, -1358, -8009, 4983, 7063, 3104, -1560, 4080, 2763, 5616, -2375, 2848, 1394, -7173, -5225, -8244, -809, 8025, -4072, -4391, -9579, 1407, 6700, 2421, -6685, 5481, -1732, -8892, -6645, 3077, 3287, -4149, 8701, -4393, -9070, -1777, 2237, -3253, -506, -4931, -7366, -8132, 5406, -6300, -275, -1908, 67, 3569, 1433, -7262, -437, 8303, 4498, -379, 3054, -6285, 4203, 6908, 4433, 3077, 2288, 9733, -8067, 3007, 9725, 9669, 1362, -2561, -4225, 5442, -9006, -429, 160, -9234, -4444, 3586, -5711, -9506, -79, -4418, -4348, -5891}
	//k := 93

	res := findMaxAverage(nums, k)
	fmt.Println(res)
}

func findMaxAverage(nums []int, k int) float64 {
	res := -1.7e+308
	l, r := 0, 0
	var sum float64

	for r < k {
		sum += float64(nums[r]) / float64(k)
		r++
	}
	res = max(res, sum)
	for r < len(nums) {
		sum = sum - (float64(nums[l]) / float64(k)) + (float64(nums[r]) / float64(k))
		res = max(res, sum)
		l++
		r++
	}

	return res
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
