package pro206

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestReverseList(t *testing.T) {
	cases := []struct {
		l        *ListNode
		expected *ListNode
	}{
		{
			l:        NewListNode([]int{1, 2, 3, 4, 5}),
			expected: NewListNode([]int{5, 4, 3, 2, 1}),
		},
		{
			l:        NewListNode([]int{1, 2}),
			expected: NewListNode([]int{2, 1}),
		},
		{
			l:        NewListNode([]int{}),
			expected: NewListNode([]int{}),
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := reverseList(tc.l)
			assert.Equal(t, tc.expected, got)
		})
	}
}
