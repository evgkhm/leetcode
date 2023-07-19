package main

import "fmt"

func main() {
	nums := make([]int, 3)
	nums[0] = 3
	nums[1] = 2
	nums[2] = 3
	//nums[2] = 4
	//nums[3] = 15

	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)

	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			if nums[j]+nums[i] == target {
				res[0], res[1] = j, i
			}
		}
	}
	return res
}
