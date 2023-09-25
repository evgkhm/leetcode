package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//l1111111 := ListNode{Val: 9, Next: nil}
	//l111111 := ListNode{Val: 9, Next: &l1111111}
	//l11111 := ListNode{Val: 9, Next: &l111111}
	//l1111 := ListNode{Val: 9, Next: &l11111}
	//l111 := ListNode{Val: 9, Next: &l1111}
	//l11 := ListNode{Val: 9, Next: &l111}
	//l1 := ListNode{Val: 9, Next: &l11}
	//
	//l2222 := ListNode{Val: 9, Next: nil}
	//l222 := ListNode{Val: 9, Next: &l2222}
	//l22 := ListNode{Val: 9, Next: &l222}
	//l2 := ListNode{Val: 9, Next: &l22}
	//l111 := ListNode{Val: 3, Next: nil}
	//l11 := ListNode{Val: 4, Next: &l111}
	//l1 := ListNode{Val: 2, Next: &l11}
	//
	//l222 := ListNode{Val: 4, Next: nil}
	//l22 := ListNode{Val: 6, Next: &l222}
	//l2 := ListNode{Val: 5, Next: &l22}

	l1 := ListNode{Val: 0, Next: nil}

	l22 := ListNode{Val: 8, Next: nil}
	l2 := ListNode{Val: 1, Next: &l22}
	res := addTwoNumbers(&l1, &l2)
	fmt.Println(res)
}

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp := dummy

	add := 0
	for l1 != nil || l2 != nil || add > 0 {
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += add
		add = 0
		if sum > 9 {
			add = 1
			sum %= 10
		}
		newNode := &ListNode{Val: sum}
		temp.Next = newNode
		temp = temp.Next
	}
	return dummy.Next
}
