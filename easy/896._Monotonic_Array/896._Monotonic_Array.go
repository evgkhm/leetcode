package main

import "fmt"

const (
	undef = 0
	incr  = 1
	decr  = 2
)

func main() {
	nums := []int{1, 1, 0}
	fmt.Println(isMonotonic(nums))
}

func isMonotonic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	state := undef

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			if state == decr {
				return false
			}
			state = incr
		} else if nums[i] < nums[i-1] {
			if state == incr {
				return false
			}
			state = decr
		}
	}

	return true
}
