package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0} //10
	k := 2
	//nums := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1} //10
	//k := 3
	//nums := []int{0, 1, 1} //2
	//k := 0
	res := longestOnes(nums, k)
	fmt.Println(res)
}

func longestOnes(nums []int, k int) int {
	l, r := 0, 0

	maxCountOnes := 0

	countFlips := 0
	for r < len(nums) {
		if nums[r] == 0 {
			countFlips++

			for countFlips > k {
				if nums[l] == 0 {
					countFlips--
				}
				l++
			}
		}

		countOnes := r - l + 1

		maxCountOnes = max(maxCountOnes, countOnes)

		r++
	}

	return maxCountOnes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
