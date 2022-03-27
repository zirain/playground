package pro2

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		nums1, nums2 []int
		expected     []int
	}{
		{
			nums1:    []int{9, 9, 9, 9, 9, 9, 9},
			nums2:    []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			l1 := NewListNode(tc.nums1)
			l2 := NewListNode(tc.nums2)
			got := addTwoNumbers(l1, l2)
			assert.Equal(t, tc.expected, got.ToInts())
		})
	}
}
