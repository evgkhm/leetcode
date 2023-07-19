package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//TEST CASE1
	/*node4R := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	node4L := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	node3R := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node3L := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	node2 := TreeNode{
		Val:   2,
		Left:  &node4L,
		Right: &node3R,
	}

	node1 := TreeNode{
		Val:   2,
		Left:  &node3L,
		Right: &node4R,
	}

	node := TreeNode{
		Val:   1,
		Left:  &node1,
		Right: &node2,
	}*/

	//TEST CASE2
	/*node3RL := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	node3RR := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	node2R := TreeNode{
		Val:   2,
		Left:  nil,
		Right: &node3RR,
	}

	node2L := TreeNode{
		Val:   2,
		Left:  nil,
		Right: &node3RL,
	}

	node := TreeNode{
		Val:   1,
		Left:  &node2L,
		Right: &node2R,
	}*/

	//TEST CASE3
	/*node2 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	node1 := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	node := TreeNode{
		Val:   1,
		Left:  &node1,
		Right: &node2,
	}*/

	//TEST CASE4
	node6 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	node5 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	node4 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	node2 := TreeNode{
		Val:   3,
		Left:  &node4,
		Right: &node5,
	}

	node3 := TreeNode{
		Val:   3,
		Left:  &node6,
		Right: nil,
	}

	node1 := TreeNode{
		Val:   2,
		Left:  &node2,
		Right: &node3,
	}

	isSymmetric(&node1)

}

func isSymmetricHelper(res *bool, rootLeft *TreeNode, rootRight *TreeNode) bool {
	if rootLeft == nil && rootRight == nil {
		return *res
	}

	if rootLeft == nil || rootRight == nil {
		*res = false
		return *res
	}

	if rootLeft.Val == rootRight.Val {
		isSymmetricHelper(res, rootLeft.Left, rootRight.Right)
		isSymmetricHelper(res, rootLeft.Right, rootRight.Left)
	} else {
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
