package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//TEST CASE1
	/*node9 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	node8 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: &node9,
	}
	node7 := TreeNode{
		Val:   13,
		Left:  nil,
		Right: nil,
	}

	node6 := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	node5 := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	node4 := TreeNode{
		Val:   11,
		Left:  &node5,
		Right: &node6,
	}

	node3 := TreeNode{
		Val:   8,
		Left:  &node7,
		Right: &node8,
	}

	node2 := TreeNode{
		Val:   4,
		Left:  &node4,
		Right: nil,
	}

	node1 := TreeNode{
		Val:   5,
		Left:  &node2,
		Right: &node3,
	}*/

	//TEST CASE2
	/*node3 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	node2 := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	node1 := TreeNode{
		Val:   1,
		Left:  &node2,
		Right: &node3,
	}*/

	//TEST CASE4
	node2 := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	node1 := TreeNode{
		Val:   1,
		Left:  &node2,
		Right: nil,
	}

	hasPathSum(&node1, 22)

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	res := false
	currSum := 0
	res = hasPathSumHelper(&res, root, targetSum, currSum)
	fmt.Println(res)
	return res
}

func hasPathSumHelper(res *bool, root *TreeNode, targetSum, currSum int) bool {
	if root == nil {
		currSum = 0
		return *res
	}

	currSum += root.Val
	if currSum == targetSum && root.Left == nil && root.Right == nil {
		*res = true
		return *res
	}

	hasPathSumHelper(res, root.Left, targetSum, currSum)
	hasPathSumHelper(res, root.Right, targetSum, currSum)

	return *res
}
