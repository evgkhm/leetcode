package main

import (
	"fmt"
	"math"
)

func main() {
	//nums := make([]int, 4)
	//nums[0] = 1
	//nums[1] = 2
	//nums[2] = 3
	//nums[3] = 1
	//k := 3 //true

	//nums := make([]int, 6)
	//nums[0] = 1
	//nums[1] = 2
	//nums[2] = 3
	//nums[3] = 1
	//nums[4] = 2
	//nums[5] = 3
	//k := 2 //false

	nums := make([]int, 4)
	nums[0] = 1
	nums[1] = 0
	nums[2] = 1
	nums[3] = 1
	k := 1 //true

	res := containsNearbyDuplicate(nums, k)

	fmt.Println(res)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if int(math.Abs(float64(i-j))) > k {
				break
			}
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
