package pro19

import (
	"github.com/zirain/leetcode/pkg/types"
)

func removeNthFromEnd(head *types.ListNode, n int) *types.ListNode {
	pre := &types.ListNode{
		Val:  -1,
		Next: head,
	}

	slow, fast := pre, pre
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return pre.Next
}
