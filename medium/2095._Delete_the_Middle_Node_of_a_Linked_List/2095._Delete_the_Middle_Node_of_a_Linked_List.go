package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node6 := ListNode{Val: 6, Next: nil}
	node5 := ListNode{Val: 2, Next: &node6}
	node4 := ListNode{Val: 1, Next: &node5}
	node3 := ListNode{Val: 7, Next: &node4}
	node2 := ListNode{Val: 4, Next: &node3}
	node1 := ListNode{Val: 3, Next: &node2}
	head := ListNode{Val: 1, Next: &node1}

	deleteMiddle(&head)
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	slow, fast := head, head
	prevSlow := &ListNode{}

	for fast.Next != nil {
		prevSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prevSlow.Next = prevSlow.Next.Next

	return head
}

//func deleteMiddle(head *ListNode) *ListNode {
//	length := 0
//	findLength(head, &length)
//	if length == 1 {
//		return nil
//	}
//	middle := (length / 2) - 1
//
//	res := head
//
//	for i := 0; i <= middle; i++ {
//		if i == middle {
//			res.Next = res.Next.Next
//		} else {
//			res = res.Next
//		}
//	}
//
//	return head
//}
//
//func findLength(head *ListNode, length *int) {
//	if head == nil {
//		return
//	}
//	*length++
//	findLength(head.Next, length)
//	return
//}
