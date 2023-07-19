package main

import "fmt"

func main() {
	//nums1 := make([]int, 5)
	//nums1[0] = 1
	//nums1[1] = 2
	//nums1[2] = 3
	//nums1[3] = 2
	//nums1[4] = 1
	//nums2 := make([]int, 5)
	//nums2[0] = 3
	//nums2[1] = 2
	//nums2[2] = 1
	//nums2[3] = 4
	//nums2[4] = 7

	//nums1 := make([]int, 5)
	//nums1[0] = 0
	//nums1[1] = 0
	//nums1[2] = 0
	//nums1[3] = 0
	//nums1[4] = 0
	//nums2 := make([]int, 5)
	//nums2[0] = 0
	//nums2[1] = 0
	//nums2[2] = 0
	//nums2[3] = 0
	//nums2[4] = 0

	//nums1 := make([]int, 5)
	//nums1[0] = 0
	//nums1[1] = 0
	//nums1[2] = 0
	//nums1[3] = 0
	//nums1[4] = 1
	//nums2 := make([]int, 5)
	//nums2[0] = 1
	//nums2[1] = 0
	//nums2[2] = 0
	//nums2[3] = 0
	//nums2[4] = 0

	//nums1 := make([]int, 5)
	//nums1[0] = 1
	//nums1[1] = 2
	//nums1[2] = 3
	//nums1[3] = 2
	//nums1[4] = 1
	//nums2 := make([]int, 4)
	//nums2[0] = 3
	//nums2[1] = 2
	//nums2[2] = 1
	//nums2[3] = 4

	nums1 := make([]int, 5)
	nums1[0] = 5
	nums1[1] = 14
	nums1[2] = 53
	nums1[3] = 80
	nums1[4] = 48
	nums2 := make([]int, 5)
	nums2[0] = 50
	nums2[1] = 47
	nums2[2] = 3
	nums2[3] = 80
	nums2[4] = 83

	res := findLength(nums1, nums2)
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLength(nums1 []int, nums2 []int) int {
	r := len(nums1) - 1
	l := 0
	res := 0

	for r >= 0 {
		for l < len(nums2) {
			counter := 0
			nums1i := r
			for nums2i := 0; nums2i <= l; nums2i++ {
				if nums1[nums1i] == nums2[nums2i] {
					counter++
					res = max(res, counter)
				} else {
					counter = 0
				}
				nums1i++
			}
			l++
			r--
		}
		r--
	}

	r = 0
	l = 0
	for l < len(nums2) {
		for r >= 0 && l < len(nums2) {
			counter := 0
			nums2i := l
			for nums1i := 0; nums1i < len(nums2)-l; nums1i++ {
				if nums1[nums1i] == nums2[nums2i] {
					counter++
					res = max(res, counter)
				} else {
					counter = 0
				}
				nums2i++
			}
			r++
			l++
		}
		l++
	}

	return res
}
