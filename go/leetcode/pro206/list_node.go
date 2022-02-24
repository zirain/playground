package pro206

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(inputs []int) *ListNode {
	if len(inputs) == 0 {
		return nil
	}

	head := &ListNode{}
	pre := head
	for _, i := range inputs {
		cur := &ListNode{
			Val: i,
		}
		pre.Next = cur
		pre = cur
	}

	return head.Next
}

func (n *ListNode) ToInts() []int {
	if n == nil {
		return []int{}
	}

	ans := make([]int, 0)

	cur := n
	for cur != nil {
		ans = append(ans, cur.Val)

		cur = cur.Next
	}
	return ans
}
