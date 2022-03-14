package pro393

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestValidUtf8(t *testing.T) {
	cases := []struct {
		data     []int
		expected bool
	}{
		{
			data:     []int{197, 130, 1},
			expected: true,
		},
		{
			data:     []int{8, 21, 212, 177, 245, 187, 178, 157, 18, 196, 155, 227, 164, 170, 5, 196, 144, 229, 187, 161, 235, 164, 188, 12, 198, 163, 245, 135, 188, 189, 113, 77, 202, 177, 213, 166, 237, 129, 180, 247, 174, 169, 167, 217, 142, 228, 154, 184, 245, 163, 146, 157, 107, 225, 182, 180, 243, 149, 146, 153, 210, 158, 30, 201, 167, 227, 160, 160},
			expected: true,
		},
		{
			data:     []int{248, 130, 130, 130},
			expected: false,
		},
		{
			data:     []int{197, 130},
			expected: true,
		},
		{
			data:     []int{235, 140, 4},
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := validUtf8(tc.data)
			assert.Equal(t, tc.expected, got)
		})
	}
}
