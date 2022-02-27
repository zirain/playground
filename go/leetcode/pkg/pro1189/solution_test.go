package pro1189

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMaxNumberOfBalloons(t *testing.T) {
	cases := []struct {
		text     string
		expected int
	}{
		{
			text:     "balloon",
			expected: 1,
		},
		{
			text:     "ballo",
			expected: 0,
		},
		{
			text:     "nlaebolko",
			expected: 1,
		},
		{
			text:     "loonbalxballpoon",
			expected: 2,
		},
		{
			text:     "leetcode",
			expected: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.text, func(t *testing.T) {
			got := maxNumberOfBalloons(tc.text)
			assert.Equal(t, tc.expected, got)
		})
	}
}
