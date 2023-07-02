package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	inorder := make([]int, 5)
	inorder[0] = 9
	inorder[1] = 3
	inorder[2] = 15
	inorder[3] = 20
	inorder[4] = 7

	postorder := make([]int, 5)
	postorder[0] = 9
	postorder[1] = 15
	postorder[2] = 7
	postorder[3] = 20
	postorder[4] = 3

	root := buildTree(inorder, postorder)
	fmt.Println(root)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 1 {
		root := TreeNode{Val: inorder[0]}
		return &root
	}
	if len(inorder) == 0 {
		return nil
	}

	// последнее число в postorder это root
	targetValPostorder := postorder[len(postorder)-1]

	// находим индекс числа targetValPostorder в inorder
	var targetIdxInorder int
	for idx, val := range inorder {
		if val == targetValPostorder {
			targetIdxInorder = idx
		}
	}

	// всё, что левее targetIdxInorder в inorder
	// и всё, что до этого индекса в postorder
	// идёт в новые inorder и postorder и в левый узел
	root := TreeNode{Val: targetValPostorder}
	root.Left = buildTree(inorder[0:targetIdxInorder], postorder[0:targetIdxInorder])

	// всё, что правее targetIdxInorder в inorder
	// и всё, что правее targetIdxInorder в postorder до targetValPostorder
	// идёт в новые inorder и postorder и в правый узел
	targetIdxPostorder := len(postorder) - 1
	root.Right = buildTree(inorder[targetIdxInorder+1:], postorder[targetIdxInorder:targetIdxPostorder])

	return &root
}
