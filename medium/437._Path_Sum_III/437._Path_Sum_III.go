package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 10, Left: nil, Right: nil}

	leftNode := &TreeNode{Val: 5, Left: nil, Right: nil}
	rightNode := &TreeNode{Val: -3, Left: nil, Right: nil}
	root.Left = leftNode
	root.Right = rightNode

	leftNode1 := &TreeNode{Val: 3, Left: nil, Right: nil}
	rightNode1 := &TreeNode{Val: 2, Left: nil, Right: nil}
	leftNode.Left = leftNode1
	leftNode.Right = rightNode1

	leftNode2 := &TreeNode{Val: 3, Left: nil, Right: nil}
	rightNode2 := &TreeNode{Val: -2, Left: nil, Right: nil}
	leftNode1.Left = leftNode2
	leftNode1.Right = rightNode2

	rightNode3 := &TreeNode{Val: 1, Left: nil, Right: nil}
	rightNode1.Right = rightNode3

	rightNode4 := &TreeNode{Val: 11, Left: nil, Right: nil}
	rightNode.Right = rightNode4

	fmt.Println(pathSum(root, 8))
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var res, sum int

	dfs(root, targetSum, &res, sum)

	return res + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func dfs(root *TreeNode, targetSum int, res *int, sum int) {
	if root == nil {
		return
	}
	sum += root.Val
	if sum == targetSum {
		*res++
	}

	dfs(root.Left, targetSum, res, sum)
	dfs(root.Right, targetSum, res, sum)
}
