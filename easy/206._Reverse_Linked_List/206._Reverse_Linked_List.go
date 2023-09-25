package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l11111 := ListNode{Val: 5, Next: nil}
	l1111 := ListNode{Val: 4, Next: &l11111}
	l111 := ListNode{Val: 3, Next: &l1111}
	l11 := ListNode{Val: 2, Next: &l111}
	l1 := ListNode{Val: 1, Next: &l11}

	res := reverseList(&l1)
	fmt.Println(res)
}

// Реверс связного списка
func reverseList(head *ListNode) *ListNode {
	var tempSlice []int
	for head != nil {
		tempSlice = append(tempSlice, head.Val)
		head = head.Next
	}

	dummy := &ListNode{}
	temp := dummy
	for i := len(tempSlice) - 1; i >= 0; i-- {
		newNode := &ListNode{Val: tempSlice[i]}
		temp.Next = newNode
		temp = temp.Next
	}
	return dummy.Next
}
