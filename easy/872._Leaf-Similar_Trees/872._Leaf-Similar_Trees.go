package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var root1ResString, root2ResString strings.Builder
	traverse(root1, &root1ResString)
	traverse(root2, &root2ResString)

	if root1ResString.String() == root2ResString.String() {
		return true
	} else {
		return false
	}
}

func traverse(root *TreeNode, rootResSting *strings.Builder) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil { // Leaf
		rootResSting.WriteString(strconv.Itoa(root.Val))
		rootResSting.WriteString(" ")
		return
	}

	traverse(root.Left, rootResSting)
	traverse(root.Right, rootResSting)
}
