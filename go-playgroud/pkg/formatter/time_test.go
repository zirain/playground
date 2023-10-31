package formatter

import (
	"fmt"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

func TestDuration(t *testing.T) {
	cases := []struct {
		d        time.Duration
		expected string
	}{
		{
			d:        time.Minute,
			expected: "1m",
		},
	}
	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := fmt.Sprintf("%d", tc.d)
			assert.Equal(t, tc.expected, got)
		})
	}
}
