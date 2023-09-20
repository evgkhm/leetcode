package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue []int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = 0 // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func main() {
	node33 := ListNode{Val: 6, Next: nil}
	node3 := ListNode{Val: 2, Next: &node33}

	node222 := ListNode{Val: 4, Next: nil}
	node22 := ListNode{Val: 3, Next: &node222}
	node2 := ListNode{Val: 1, Next: &node22}

	node111 := ListNode{Val: 5, Next: nil}
	node11 := ListNode{Val: 4, Next: &node111}
	node1 := ListNode{Val: 1, Next: &node11}

	var sliceOfNodes []*ListNode
	sliceOfNodes = append(sliceOfNodes, &node1)
	sliceOfNodes = append(sliceOfNodes, &node2)
	sliceOfNodes = append(sliceOfNodes, &node3)

	mergeKLists(sliceOfNodes)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}

	var prQ PriorityQueue

	for _, node := range lists {
		for node != nil {
			heap.Push(&prQ, node.Val)
			node = node.Next
		}
	}

	dummy := &ListNode{}
	tmp := dummy
	for prQ.Len() > 0 {
		value := heap.Pop(&prQ).(int)
		newNode := &ListNode{Val: value}
		tmp.Next = newNode
		tmp = tmp.Next
	}
	return dummy.Next
}
