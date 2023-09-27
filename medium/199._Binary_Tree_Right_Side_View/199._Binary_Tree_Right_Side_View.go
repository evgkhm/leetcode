package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1, Left: nil, Right: nil}
	l1 := &TreeNode{Val: 2, Left: nil, Right: nil}
	r1 := &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left, root.Right = l1, r1
	l11 := &TreeNode{Val: 5, Left: nil, Right: nil}
	r11 := &TreeNode{Val: 4, Left: nil, Right: nil}
	l1.Left, r1.Right = l11, r11

	var res []int
	res = rightSideView(root)
	fmt.Println(res)
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)

		currNode := queue[0]
		queue = queue[1:]
		res = append(res, currNode.Val)

		if currNode.Right != nil {
			queue = append(queue, currNode.Right)
		}
		if currNode.Left != nil {
			queue = append(queue, currNode.Left)
		}

		for i := 1; i < size; i++ {
			tmp := queue[0]
			queue = queue[1:]

			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
		}
	}

	return res
}
