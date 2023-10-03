package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedNums := mergeSort(nums1, nums2)

	if len(mergedNums)%2 == 0 {
		mid1Idx := (len(mergedNums) / 2) - 1
		mid2Idx := len(mergedNums) / 2

		mid1Val := float64(mergedNums[mid1Idx])
		mid2Val := float64(mergedNums[mid2Idx])

		return float64((mid1Val + mid2Val) / 2)
	} else {
		return float64(mergedNums[len(mergedNums)/2])
	}
}

func mergeSort(nums1 []int, nums2 []int) []int {
	var res []int
	p1, p2 := 0, 0

	for p1+p2 < len(nums1)+len(nums2) {
		if p1 < len(nums1) && p2 < len(nums2) {
			if nums1[p1] < nums2[p2] {
				res = append(res, nums1[p1])
				p1++
			} else {
				res = append(res, nums2[p2])
				p2++
			}
		} else if p1 < len(nums1) {
			res = append(res, nums1[p1])
			p1++
		} else {
			res = append(res, nums2[p2])
			p2++
		}
	}

	return res
}
