package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//nums := []int{1, 2, 3, 5, 6} // 1
	//nums := []int{4, 2, 5, 3} // 0
	//nums := []int{1, 10, 100, 1000} // 3
	//nums := []int{8, 5, 9, 9, 8, 4} // 2
	//nums := []int{41, 33, 29, 33, 35, 26, 47, 24, 18, 28} // 5
	nums := []int{8, 10, 16, 18, 10, 10, 16, 13, 13, 16} // 6
	fmt.Println(minOperations(nums))
}

func minOperations(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	valuesMap := make(map[int]bool)

	for _, val := range nums {
		valuesMap[val] = true
	}

	var sortedUnduplicatedNums []int
	for key, _ := range valuesMap {
		sortedUnduplicatedNums = append(sortedUnduplicatedNums, key)
	}

	sort.Slice(sortedUnduplicatedNums, func(i, j int) bool {
		return sortedUnduplicatedNums[i] < sortedUnduplicatedNums[j]
	})

	resCounter := math.MaxInt

	j := 0
	for i := 0; i < len(sortedUnduplicatedNums); i++ {
		for j < len(sortedUnduplicatedNums) && sortedUnduplicatedNums[j] < sortedUnduplicatedNums[i]+len(nums) {
			j += 1
		}
		counter := j - i
		resCounter = min(resCounter, len(nums)-counter)
	}

	return resCounter
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
