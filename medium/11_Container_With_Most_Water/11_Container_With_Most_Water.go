package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//height := []int{2, 3, 4, 5, 18, 17, 6}
	res := maxArea(height)
	fmt.Println(res)
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0

	for l < r {
		var currArea int
		if height[l] > height[r] {
			currArea = (r - l) * height[r]
			r--
		} else {
			currArea = (r - l) * height[l]
			l++
		}

		res = max(res, currArea)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
