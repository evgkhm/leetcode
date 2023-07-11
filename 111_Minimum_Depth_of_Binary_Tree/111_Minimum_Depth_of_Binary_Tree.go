package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//node4 := TreeNode{Val: 7, Left: nil, Right: nil}
	//node3 := TreeNode{Val: 15, Left: nil, Right: nil}
	//node2 := TreeNode{Val: 20, Left: &node3, Right: &node4}
	//node1 := TreeNode{Val: 9, Left: nil, Right: nil}
	//tree := TreeNode{Val: 3, Left: &node1, Right: &node2}

	node4 := TreeNode{Val: 6, Left: nil, Right: nil}
	node3 := TreeNode{Val: 5, Left: nil, Right: &node4}
	node2 := TreeNode{Val: 4, Left: nil, Right: &node3}
	node1 := TreeNode{Val: 3, Left: nil, Right: &node2}
	tree := TreeNode{Val: 2, Left: nil, Right: &node1}

	res := minDepth(&tree)
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return 1 + dfs(root.Right)
	} else if root.Right == nil {
		return 1 + dfs(root.Left)
	}

	return 1 + min(dfs(root.Left), dfs(root.Right))
}

func minDepth(root *TreeNode) int {
	res := 0
	res = dfs(root)
	return res
}
