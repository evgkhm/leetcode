package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node6 := ListNode{Val: 7, Next: nil}
	node5 := ListNode{Val: 4, Next: &node6}
	node4 := ListNode{Val: 6, Next: &node5}
	node3 := ListNode{Val: 5, Next: &node4}
	node2 := ListNode{Val: 3, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	head := ListNode{Val: 2, Next: &node1}

	//node4 := ListNode{Val: 5, Next: nil}
	//node3 := ListNode{Val: 4, Next: &node4}
	//node2 := ListNode{Val: 3, Next: &node3}
	//node1 := ListNode{Val: 2, Next: &node2}
	//head := ListNode{Val: 1, Next: &node1}

	oddEvenList(&head)
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	oddTail, evenTail := head, head.Next
	evenHead := evenTail

	for evenTail != nil && evenTail.Next != nil {
		oddTail.Next = evenTail.Next
		oddTail = oddTail.Next
		evenTail.Next = oddTail.Next
		evenTail = evenTail.Next
	}
	oddTail.Next = evenHead
	return head

}
