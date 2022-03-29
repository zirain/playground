package pro2024

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMaxConsecutiveAnswers(t *testing.T) {
	cases := []struct {
		answerKey string
		k         int
		expected  int
	}{
		{
			answerKey: "TTFF",
			k:         2,
			expected:  4,
		},
		{
			answerKey: "TFFT",
			k:         1,
			expected:  3,
		},
		{
			answerKey: "TTFTTFTT",
			k:         1,
			expected:  5,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := maxConsecutiveAnswers(tc.answerKey, tc.k)
			assert.Equal(t, tc.expected, got)
		})
	}
}
