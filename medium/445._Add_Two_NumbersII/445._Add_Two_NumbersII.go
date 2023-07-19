package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//l222 := ListNode{Val: 4, Next: nil}
	//l22 := ListNode{Val: 6, Next: &l222}
	//l2 := ListNode{Val: 5, Next: &l22}
	//
	//l1111 := ListNode{Val: 3, Next: nil}
	//l111 := ListNode{Val: 4, Next: &l1111}
	//l11 := ListNode{Val: 2, Next: &l111}
	//l1 := ListNode{Val: 7, Next: &l11} // 7 8 0 7

	//l2 := ListNode{Val: 5, Next: nil}
	//l1 := ListNode{Val: 5, Next: nil} // 1 0

	//TEST1062
	l22 := ListNode{Val: 9, Next: nil}
	l2 := ListNode{Val: 9, Next: &l22}
	l1 := ListNode{Val: 1, Next: nil} // 1 1 0

	addTwoNumbers(&l1, &l2)
}

func builtSliceFromListNode(l *ListNode, lSlice *[]int) {
	if l == nil {
		return
	}
	*lSlice = append(*lSlice, l.Val)
	builtSliceFromListNode(l.Next, lSlice)
	return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1Slice []int
	var l2Slice []int
	builtSliceFromListNode(l1, &l1Slice)
	builtSliceFromListNode(l2, &l2Slice)

	lenL1 := len(l1Slice)
	lenL2 := len(l2Slice)

	resLenSlice := 0
	number, one := 0, 0
	if lenL1 > lenL2 {
		resLenSlice = lenL1
	} else {
		resLenSlice = lenL2
	}
	resSlice := make([]int, resLenSlice)
	idx := 0

	for idx < resLenSlice {
		if lenL1 >= 0 {
			lenL1--
		}
		if lenL2 >= 0 {
			lenL2--
		}

		if lenL1 >= 0 && lenL2 >= 0 {
			number = l1Slice[lenL1] + l2Slice[lenL2] + resSlice[idx]
		} else if lenL1 >= 0 {
			number = l1Slice[lenL1] + resSlice[idx]
		} else if lenL2 >= 0 {
			number = l2Slice[lenL2] + resSlice[idx]
		}
		one = number / 10
		number %= 10
		resSlice[idx] = number

		idx++
		if one > 0 {
			if len(resSlice) <= idx { //if it's end of slice (node)
				resSlice = append(resSlice, one)
				continue
			}
			resSlice[idx] = one
		}
	}

	lenSlice := len(resSlice) - 1
	res := createListNodeFromSlice(resSlice, &lenSlice)
	return res
}

func createListNodeFromSlice(slice []int, lenSlice *int) *ListNode {
	if *lenSlice < 0 {
		return nil
	}

	node := ListNode{}
	node.Val = slice[*lenSlice]
	*lenSlice--
	node.Next = createListNodeFromSlice(slice, lenSlice)

	return &node
}
