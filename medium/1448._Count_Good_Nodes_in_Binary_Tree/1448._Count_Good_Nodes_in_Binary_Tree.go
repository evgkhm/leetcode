package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	rightNode1 := &TreeNode{Val: 4, Left: nil, Right: nil}
	leftNode1 := &TreeNode{Val: 4, Left: nil, Right: nil}

	leftNode := &TreeNode{Val: 3, Left: leftNode1, Right: rightNode1}

	root := &TreeNode{Val: 3, Left: leftNode, Right: nil}

	fmt.Println(goodNodes(root))
}

func goodNodes(root *TreeNode) int {
	res := 1
	dfs(root.Left, &res, root.Val)
	dfs(root.Right, &res, root.Val)
	return res
}

func dfs(root *TreeNode, res *int, maxVal int) {
	if root == nil {
		return
	}
	if root.Val >= maxVal {
		*res++
		maxVal = root.Val
	}

	dfs(root.Left, res, maxVal)
	dfs(root.Right, res, maxVal)
}
