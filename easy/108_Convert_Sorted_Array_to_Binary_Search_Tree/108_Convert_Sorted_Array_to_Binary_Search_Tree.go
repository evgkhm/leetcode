package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := make([]int, 5)
	nums[0] = -10
	nums[1] = -3
	nums[2] = 0
	nums[3] = 5
	nums[4] = 9
	sortedArrayToBST(nums)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	if len(nums) == 0 {
		return nil
	}
	var centerIdx int
	centerIdx = int(math.Floor(float64(len(nums) / 2)))
	root := TreeNode{Val: nums[centerIdx]}

	root.Left = sortedArrayToBST(nums[0:centerIdx])
	root.Right = sortedArrayToBST(nums[centerIdx+1:])

	return &root
}
