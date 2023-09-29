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
			if state == undef {
				state = incr
			} else if state == decr {
				return false
			}
		} else if nums[i] < nums[i-1] {
			if state == undef {
				state = decr
			} else if state == incr {
				return false
			}
		}
	}

	return true
}
