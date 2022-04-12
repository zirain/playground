package pro429

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestLevel(t *testing.T) {

	cases := []struct {
		root     *Node
		expected [][]int
	}{
		{
			root: &Node{
				Val: 1,
				Children: []*Node{
					{
						Val: 3,
						Children: []*Node{
							{Val: 5},
							{Val: 6},
						},
					},
					{
						Val: 2,
					},
					{
						Val: 4,
					},
				},
			},
			expected: [][]int{{1}, {3, 2, 4}, {5, 6}},
		},
		{
			root: &Node{
				Val: 1,
				Children: []*Node{
					{
						Val: 2,
					},
					{
						Val: 3,
						Children: []*Node{
							{Val: 6},
							{
								Val: 7,
								Children: []*Node{
									{
										Val: 11,
										Children: []*Node{
											{Val: 14},
										},
									},
								},
							},
						},
					},
					{
						Val: 4,
						Children: []*Node{
							{
								Val: 8,
								Children: []*Node{
									{Val: 12},
								},
							},
						},
					},
					{
						Val: 5,
						Children: []*Node{
							{
								Val: 9,
								Children: []*Node{
									{Val: 13},
								},
							},
							{Val: 10},
						},
					},
				},
			},
			expected: [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := levelOrder(tc.root)
			assert.Equal(t, tc.expected, got)
		})
	}
}
