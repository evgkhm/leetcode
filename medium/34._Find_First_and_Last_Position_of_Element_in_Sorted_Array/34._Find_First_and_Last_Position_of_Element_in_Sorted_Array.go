package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8

	fmt.Println(searchRange(nums, target))
}

// Time: O(log n)
func searchRangeLeft(nums []int, target int) int {
	l, r := 0, len(nums)-1
	idx := -1

	for l <= r {
		mid := (r + l) / 2

		if nums[mid] == target {
			idx = mid
			r = mid - 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return idx
}

func searchRangeRight(nums []int, target int) int {
	l, r := 0, len(nums)-1
	idx := -1

	for l <= r {
		mid := (r + l) / 2

		if nums[mid] == target {
			idx = mid
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return idx
}

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	l = searchRangeLeft(nums, target)
	r = searchRangeRight(nums, target)
	if l <= r {
		return []int{l, r}
	} else {
		return []int{-1, -1}
	}
}

// Time : O(n)
//func searchRange(nums []int, target int) []int {
//	l, r := -1, -1
//	start := false
//	for indx, val := range nums {
//		if val == target && start == false {
//			l = indx
//			start = true
//		}
//		if val == target && start == true {
//			r = indx
//		}
//	}
//	return []int{l, r}
//}
