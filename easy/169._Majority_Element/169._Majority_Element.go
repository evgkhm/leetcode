package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	counter := 1
	majorityElem := nums[0]

	for _, val := range nums[1:] {
		if val == majorityElem {
			counter += 1
		} else {
			counter -= 1
		}

		if counter == 0 {
			majorityElem = val
			counter = 1
		}
	}
	return majorityElem
}

// Second version
//func majorityElement(nums []int) int {
//	if len(nums) == 1 {
//		return nums[0]
//	}
//
//	counter := 0
//	majorityElem := 0
//	for _, val := range nums {
//		if counter == 0 {
//			majorityElem = val
//			counter = 1
//		} else if val == majorityElem {
//			counter += 1
//		} else {
//			counter -= 1
//		}
//	}
//	return majorityElem
//}
