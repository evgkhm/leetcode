package main

import (
	"fmt"
	"math"
)

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
	//tree := TreeNode{Val: 3, Left: &node1, Right: &node2} //true

	//node4 := TreeNode{Val: 6, Left: nil, Right: nil}
	//node3 := TreeNode{Val: 5, Left: nil, Right: &node4}
	//node2 := TreeNode{Val: 4, Left: nil, Right: &node3}
	//node1 := TreeNode{Val: 3, Left: nil, Right: &node2}
	//tree := TreeNode{Val: 2, Left: nil, Right: &node1} //false

	//node2 := TreeNode{Val: 3, Left: nil, Right: nil}
	//node1 := TreeNode{Val: 2, Left: nil, Right: &node2}
	//tree := TreeNode{Val: 1, Left: nil, Right: &node1} //false

	node6 := TreeNode{Val: 6, Left: nil, Right: nil}
	node5 := TreeNode{Val: 8, Left: nil, Right: nil}
	node4 := TreeNode{Val: 5, Left: nil, Right: nil}
	node3 := TreeNode{Val: 4, Left: &node5, Right: nil}
	node2 := TreeNode{Val: 3, Left: nil, Right: &node6}
	node1 := TreeNode{Val: 2, Left: &node3, Right: &node4}
	tree := TreeNode{Val: 1, Left: &node1, Right: &node2} //true

	res := isBalanced(&tree)
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalancedHelper(root *TreeNode, res *bool, depth int) int {
	if root == nil {
		return 0
	}

	left := isBalancedHelper(root.Left, res, depth)
	right := isBalancedHelper(root.Right, res, depth)

	if math.Abs(float64(left-right)) > 1 {
		*res = false
	}

	return max(left, right) + 1
}

func isBalanced(root *TreeNode) bool {
	res := true
	isBalancedHelper(root, &res, 0)
	return res
}
