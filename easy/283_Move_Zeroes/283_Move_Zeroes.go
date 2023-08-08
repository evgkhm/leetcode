package main

import "fmt"

func main() {
	nums := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	l, r := 0, 1

	for l < len(nums) {
		if nums[l] == 0 {
			for r < len(nums) {
				if nums[r] != 0 {
					nums[l], nums[r] = nums[r], nums[l]
					break
				}
				r++
			}
		}

		l++
		r++
	}
}
