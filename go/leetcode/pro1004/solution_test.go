package pro1004

import "testing"

func TestLongestOnes(t *testing.T) {
	type sampleData struct {
		A      []int
		K      int
		Expect int
	}

	samples := []sampleData{
		{
			A:      []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			K:      2,
			Expect: 6,
		}, {
			A:      []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			K:      3,
			Expect: 10,
		},
	}

	for _, d := range samples {
		actual := longestOnes(d.A, d.K)
		if actual != d.Expect {
			t.Fail()
		}
	}
}
