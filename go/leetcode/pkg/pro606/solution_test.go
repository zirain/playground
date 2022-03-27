package pro606

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestTree2str(t *testing.T) {
	cases := []struct {
		tree     *TreeNode
		expected string
	}{
		{
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: "1(2(4))(3)",
		},
		{
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
				},
			},
			expected: "1(2)(3(4))",
		},
		{
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			expected: "1(2)(3()(4))",
		},
		{
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: "1(2()(4))(3)",
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := tree2str(tc.tree)
			assert.Equal(t, tc.expected, got)
		})
	}
}
