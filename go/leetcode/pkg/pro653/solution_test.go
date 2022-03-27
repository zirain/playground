package pro653

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestFindTarget(t *testing.T) {
	cases := []struct {
		root     *TreeNode
		k        int
		expected bool
	}{
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 6,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			k:        9,
			expected: true,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := findTarget(tc.root, tc.k)
			assert.Equal(t, tc.expected, got)
		})
	}
}
