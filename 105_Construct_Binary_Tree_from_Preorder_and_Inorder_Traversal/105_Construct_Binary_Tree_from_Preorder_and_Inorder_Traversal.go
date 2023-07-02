package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preorder := make([]int, 5)
	preorder[0] = 3
	preorder[1] = 9
	preorder[2] = 20
	preorder[3] = 15
	preorder[4] = 7

	inorder := make([]int, 5)
	inorder[0] = 9
	inorder[1] = 3
	inorder[2] = 15
	inorder[3] = 20
	inorder[4] = 7
	root := buildTree(preorder, inorder)
	fmt.Println(root)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	//проверки
	if len(preorder) == 1 {
		root := TreeNode{Val: preorder[0]}
		return &root
	}
	if len(preorder) == 0 {
		return nil
	}

	var targetVal, targetIdx int
	targetVal = preorder[0]
	// найти нужный индекс (targetIdx в inorder по значению targetVal из preorder)
	for idx, val := range inorder {
		if val == targetVal {
			targetIdx = idx
		}
	}

	root := TreeNode{Val: targetVal}
	//newPreorderLeft := preorder[1 : targetIdx+1]
	//newInorderLeft := inorder[0:targetIdx]
	root.Left = buildTree(preorder[1:targetIdx+1], inorder[0:targetIdx])

	//newPreorderRight := preorder[targetIdx+1:]
	//newInorderRight := inorder[targetIdx+1:]
	root.Right = buildTree(preorder[targetIdx+1:], inorder[targetIdx+1:])

	return &root
}
