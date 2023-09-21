package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node3 := ListNode{Val: 3, Next: nil}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 2, Next: &node2}
	head := ListNode{Val: 4, Next: &node1}

	fmt.Println(pairSum(&head))
}

func pairSum(head *ListNode) int {
	var sliceOfVal []int

	for head != nil {
		sliceOfVal = append(sliceOfVal, head.Val)
		head = head.Next
	}

	maxSum := 0
	for i := 0; i < len(sliceOfVal)/2; i++ {
		maxSum = max(maxSum, sliceOfVal[i]+sliceOfVal[len(sliceOfVal)-1-i])
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
