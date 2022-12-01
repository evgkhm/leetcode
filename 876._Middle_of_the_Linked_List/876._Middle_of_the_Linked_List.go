package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Print() {
	ptr := n
	for ptr != nil {
		fmt.Printf("%v\n", ptr.Val)
		ptr = ptr.Next
	}
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
	/*res := []*ListNode{}
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	return res[len(res)/2]*/
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n1.Next = n2
	n3 := &ListNode{Val: 3}
	n2.Next = n3
	n4 := &ListNode{Val: 4}
	n3.Next = n4
	n5 := &ListNode{Val: 5}
	n4.Next = n5
	n6 := &ListNode{Val: 6}
	n5.Next = n6
	n7 := &ListNode{Val: 7}
	n6.Next = n7

	//n1.Print()
	n1 = middleNode(n1)
	n1.Print()
}
