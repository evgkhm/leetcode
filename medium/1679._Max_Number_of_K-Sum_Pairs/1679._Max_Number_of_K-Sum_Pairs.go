package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 1, 3, 4, 3}
	//nums := []int{1, 2, 3, 4}
	k := 5
	res := maxOperations(nums, k)
	fmt.Println(res)
}

func maxOperations(nums []int, k int) int {
	l, r := 0, len(nums)-1
	res := 0

	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j]
	})

	for l < r {
		if nums[l]+nums[r] == k {
			nums = append(nums[:r], nums[r+1:]...) // delete r
			nums = append(nums[:l], nums[l+1:]...) // delete l
			res++
			l, r = 0, len(nums)-1
			continue
		}

		if nums[l]+nums[r] > k {
			r--
		} else {
			l++
		}
	}
	return res
}
