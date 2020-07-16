package pro35

import "testing"

func Test1(t *testing.T) {
	datas := []struct {
		nums   []int
		target int
		except int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1, 3, 5, 6}, 6, 3},
	}

	for _, data := range datas {
		actual := searchInsert(data.nums, data.target)
		if actual != data.except {
			t.Errorf("input: %v, target: %d, except: %d, actual: %d", data.nums, data.target, data.except, actual)
		}
	}
}
