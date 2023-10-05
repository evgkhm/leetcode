package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 1, 1, 3, 1, 4, 5, 6}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) []int {
	count1, count2 := 0, 0
	majorityElem1, majorityElem2 := math.MinInt32, math.MinInt32

	for _, val := range nums {
		if val == majorityElem1 {
			count1 += 1
		} else if val == majorityElem2 {
			count2 += 1
		} else if count1 == 0 {
			majorityElem1 = val
			count1 = 1
		} else if count2 == 0 {
			majorityElem2 = val
			count2 = 1
		} else {
			count1 -= 1
			count2 -= 1
		}
	}

	count1, count2 = 0, 0
	for _, val := range nums {
		if val == majorityElem1 {
			count1++
		} else if val == majorityElem2 {
			count2++
		}
	}

	var res []int
	if count1 > len(nums)/3 {
		res = append(res, majorityElem1)
	}
	if count2 > len(nums)/3 {
		res = append(res, majorityElem2)
	}

	return res
}
