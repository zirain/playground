package pro728

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

func TestSelfDividingNumbers(t *testing.T) {
	cases := []struct {
		left, right int
		expected    []int
	}{
		{
			1, 22,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
		{
			47, 85,
			[]int{48, 55, 66, 77},
		},
		{
			3056,
			4813,
			[]int{3111, 3126, 3132, 3135, 3144, 3162, 3168, 3171, 3195, 3216, 3222, 3264, 3276, 3288, 3312, 3315, 3324, 3333, 3336, 3339, 3366, 3384, 3393, 3432, 3444, 3492, 3555, 3612, 3624, 3636, 3648, 3666, 3717, 3816, 3864, 3888, 3915, 3924, 3933, 3996, 4112, 4116, 4124, 4128, 4144, 4164, 4172, 4184, 4212, 4224, 4236, 4244, 4248, 4288, 4332, 4344, 4368, 4392, 4412, 4416, 4424, 4444, 4448, 4464, 4488, 4632, 4644},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("[%d:%d]", tc.left, tc.right), func(t *testing.T) {
			got := selfDividingNumbers(tc.left, tc.right)
			assert.Equal(t, tc.expected, got)
		})
	}
}
