package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list222 := ListNode{Val: 4, Next: nil}
	list22 := ListNode{Val: 3, Next: &list222}
	list2 := ListNode{Val: 1, Next: &list22}

	list111 := ListNode{Val: 4, Next: nil}
	list11 := ListNode{Val: 2, Next: &list111}
	list1 := ListNode{Val: 1, Next: &list11}

	mergeTwoLists(&list1, &list2)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tmp := dummy
	for list1 != nil || list2 != nil {
		val1, val2 := 101, 101 // More than max
		if list1 != nil {
			val1 = list1.Val
		}
		if list2 != nil {
			val2 = list2.Val
		}
		if val1 < val2 {
			newNode := &ListNode{Val: val1}
			tmp.Next = newNode
			tmp = tmp.Next
			list1 = list1.Next
		} else {
			newNode := &ListNode{Val: val2}
			tmp.Next = newNode
			tmp = tmp.Next
			list2 = list2.Next
		}
	}
	return dummy.Next
}
