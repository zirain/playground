package pro19

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/zirain/leetcode/pkg/types"
)

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		inputs   []int
		n        int
		expected []int
	}{
		{
			inputs:   []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3, 5},
		},
		{
			inputs:   []int{1},
			n:        1,
			expected: []int{},
		},
		{
			inputs:   []int{1, 2},
			n:        1,
			expected: []int{1},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			ln := types.NewListNode(tc.inputs)
			got := removeNthFromEnd(ln, tc.n)
			assert.Equal(t, tc.expected, got.ToInts())
		})
	}
}
