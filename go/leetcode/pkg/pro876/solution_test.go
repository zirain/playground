package pro876

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/zirain/leetcode/types"
)

func TestMiddleNode(t *testing.T) {
	cases := []struct {
		inputs   []int
		expected []int
	}{
		{
			inputs:   []int{1, 2, 3, 4, 5},
			expected: []int{3, 4, 5},
		},
		{
			inputs:   []int{1, 2, 3, 4, 5, 6},
			expected: []int{4, 5, 6},
		},
		{
			inputs:   []int{1},
			expected: []int{1},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			n := types.NewListNode(tc.inputs)
			got := middleNode(n)
			assert.Equal(t, tc.expected, got.ToInts())
		})
	}
}
