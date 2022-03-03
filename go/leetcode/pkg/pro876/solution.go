package pro876

import "github.com/zirain/leetcode/pkg/types"

func middleNode(head *types.ListNode) *types.ListNode {
	pre := &types.ListNode{
		Next: head,
	}
	slow, fast := pre, pre
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow.Next
}
