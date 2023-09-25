package main

import "fmt"

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

	levelOrder(&node)
}

// Определить уровни поддеревьев
func levelHelper(res *[][]int, root *TreeNode, level int) {
	if root == nil {
		return
	}

	if level >= len(*res) {
		*res = append(*res, []int{})
	}

	(*res)[level] = append((*res)[level], root.Val)

	levelHelper(res, root.Left, level+1)
	levelHelper(res, root.Right, level+1)
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	levelHelper(&res, root, 0)
	fmt.Println(res)
	return res
}
