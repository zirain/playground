package pro0201

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	dict := make(map[int]bool)
	pre := &ListNode{
		Val:  -1,
		Next: head,
	}
	cur := head
	for cur != nil {
		if _, ok := dict[cur.Val]; ok {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			dict[cur.Val] = true
			cur = cur.Next
			pre = cur
		}
	}

	return head
}
