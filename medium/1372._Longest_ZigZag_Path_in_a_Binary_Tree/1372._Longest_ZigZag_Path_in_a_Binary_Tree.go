package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1, Left: nil, Right: nil}

	leftNode := &TreeNode{Val: 1, Left: nil, Right: nil}
	rightNode := &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = leftNode
	root.Right = rightNode

	rightNode1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	leftNode.Right = rightNode1

	rightNode3 := &TreeNode{Val: 1, Left: nil, Right: nil}
	leftNode3 := &TreeNode{Val: 1, Left: nil, Right: nil}
	rightNode1.Left = leftNode3
	rightNode1.Right = rightNode3

	rightNode4 := &TreeNode{Val: 1, Left: nil, Right: nil}
	leftNode3.Right = rightNode4

	fmt.Println(longestZigZag(root))
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := dfsLeft(root.Left)
	r := dfsRight(root.Right)
	return max(max(max(l, r), longestZigZag(root.Left)), longestZigZag(root.Right))
}

func dfsLeft(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + dfsRight(root.Right)
}

func dfsRight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + dfsLeft(root.Left)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
