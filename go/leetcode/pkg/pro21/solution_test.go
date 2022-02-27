package pro21

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		l1, l2 *ListNode

		expected *ListNode
	}{
		{
			l1:       NewListNode([]int{1, 2, 4}),
			l2:       NewListNode([]int{1, 3, 4}),
			expected: NewListNode([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			l1:       NewListNode([]int{}),
			l2:       NewListNode([]int{}),
			expected: NewListNode([]int{}),
		},
		{
			l1:       NewListNode([]int{}),
			l2:       NewListNode([]int{0}),
			expected: NewListNode([]int{0}),
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := mergeTwoLists(tc.l1, tc.l2)
			assert.Equal(t, tc.expected, got)
		})
	}
}
