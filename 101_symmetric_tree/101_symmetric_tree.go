package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node4 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	node3 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	node2 := TreeNode{
		Val:   2,
		Left:  &node4,
		Right: &node3,
	}

	node1 := TreeNode{
		Val:   2,
		Left:  &node3,
		Right: &node4,
	}

	node := TreeNode{
		Val:   1,
		Left:  &node1,
		Right: &node2,
	}

	/*node2 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	node1 := TreeNode{
		Val:   2,
		Left:  nil,
		Right: &node2,
	}

	node := TreeNode{
		Val:   1,
		Left:  &node1,
		Right: &node1,
	}*/

	isSymmetric(&node)

}

func isSymmetricHelper(res *bool, rootLeft *TreeNode, rootRight *TreeNode) bool {
	if rootLeft == nil || rootRight == nil {
		return *res
	}

	if rootLeft.Val == rootRight.Val &&
		rootLeft.Left != nil && rootLeft.Right != nil &&
		rootRight.Left != nil && rootRight.Right != nil {
		*res = true
		isSymmetricHelper(res, rootLeft.Left, rootRight.Left)
		isSymmetricHelper(res, rootLeft.Right, rootRight.Right)
	} else if rootLeft.Val != rootRight.Val {
		*res = false
		return *res
	} else if rootLeft.Left == nil && rootLeft.Right != nil {
		*res = false
		return *res
	} else if rootLeft.Left != nil && rootLeft.Right == nil {
		*res = false
		return *res
	} else if rootRight.Left == nil && rootRight.Right != nil {
		*res = false
		return *res
	} else if rootRight.Left != nil && rootRight.Right == nil {
		*res = false
		return *res
	}

	return *res
}

func isSymmetric(root *TreeNode) bool {
	res := true
	isSymmetricHelper(&res, root.Left, root.Right)
	fmt.Println(res)
	return res
}
