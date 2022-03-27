package pro2

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addListNode(l1, l2, 0)
}

func addListNode(l1 *ListNode, l2 *ListNode, addition int) *ListNode {
	if l1 == nil && l2 == nil && addition == 0 {
		return nil
	} else if l1 == nil && l2 == nil {
		return &ListNode{
			Val: addition,
		}
	}

	v, add := addition, 0
	var lNext, rNext *ListNode
	if l1 != nil {
		v += l1.Val
		lNext = l1.Next
	}
	if l2 != nil {
		v += l2.Val
		rNext = l2.Next
	}

	for v >= 10 {
		v -= 10
		add++
	}

	return &ListNode{
		Val:  v,
		Next: addListNode(lNext, rNext, add),
	}
}
