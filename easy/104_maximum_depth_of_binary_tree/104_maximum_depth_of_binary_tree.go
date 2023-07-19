package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node4 := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	node3 := TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}

	node2 := TreeNode{
		Val:   20,
		Left:  &node3,
		Right: &node4,
	}

	node1 := TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}

	node := TreeNode{
		Val:   3,
		Left:  &node1,
		Right: &node2,
	}

	/*node1 := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	node := TreeNode{
		Val:   1,
		Left:  nil,
		Right: &node1,
	}*/

	maxDepth(&node)

}

func maxDepthHelper(res *int, root *TreeNode, depth int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil { // leaf
		if depth > *res {
			*res = depth // update result
		}
		depth = 0 // reset depth to zero cause it's leaf
	}

	maxDepthHelper(res, root.Left, depth+1)
	maxDepthHelper(res, root.Right, depth+1)
}

func maxDepth(root *TreeNode) int {
	res, depth := 0, 1
	maxDepthHelper(&res, root, depth)
	fmt.Println(res)
	return res
}
