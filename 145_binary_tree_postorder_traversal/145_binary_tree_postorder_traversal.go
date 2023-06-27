package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node3 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	node2 := TreeNode{
		Val:   2,
		Left:  &node3,
		Right: nil,
	}

	node1 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: &node2,
	}

	postorderTraversal(&node1)
}

func dfs(root *TreeNode, slice *[]int) []int {
	if root == nil {
		return nil
	}
	dfs(root.Left, slice)
	dfs(root.Right, slice)
	*slice = append(*slice, root.Val)
	return *slice
}

func postorderTraversal(root *TreeNode) []int {
	var slice []int
	slice = dfs(root, &slice)
	fmt.Println(slice)
	return slice
}
