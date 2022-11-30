package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

// LinkedList Связанный список
type LinkedList struct {
	head *ListNode // Головной узел
	len  int
}

func middleNode(head *ListNode) *ListNode {
	length := 0

	p := head.Next
	for p != nil {
		p = p.Next
		length++
	}

	middle := 0
	if length&1 == 0 { //четное
		middle = (length >> 1) + 1 //середина
	} else { //нечетное
		middle = length >> 1 //середина
	}

	res := head.Next
	for p != nil && middle <= length {
		res.Next = head.Next
		middle++
	}

	return res
}

func (l *LinkedList) Insert(val int) {
	n := ListNode{}
	n.Val = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func CreateLinkedList() *LinkedList {
	head := new(ListNode)
	return &LinkedList{head}
}

func main() {
	list := CreateLinkedList()
	s := []int{1, 2, 3, 4, 5, 6, 7}

	for i, v := range s {
		list.Insert(i+1, v)
	}
	fmt.Println(list)
	res := middleNode(list)
	fmt.Println(res)
}
