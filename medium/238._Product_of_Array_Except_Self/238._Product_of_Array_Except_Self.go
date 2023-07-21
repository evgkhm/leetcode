package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	res := productExceptSelf(nums)
	fmt.Println(res)
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	pref := make([]int, n)
	suf := make([]int, n)

	pref[0] = nums[0]
	suf[n-1] = nums[n-1]

	for i := 1; i < len(nums); i++ {
		pref[i] = pref[i-1] * nums[i]
	}

	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i]
	}

	res[0] = suf[1] // first
	for i := 1; i < n-1; i++ {
		res[i] = pref[i-1] * suf[i+1]
	}
	res[n-1] = pref[n-2] // last

	return res
}
