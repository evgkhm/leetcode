package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

//func removeDuplicates(nums []int) int {
//	idx := 1
//	for idx < len(nums) {
//		if nums[idx] == nums[idx-1] {
//			nums = append(nums[:idx], nums[idx+1:]...)
//		} else {
//			idx++
//		}
//	}
//
//	return len(nums)
//}
