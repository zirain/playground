package pro954

import (
	"sort"
)

func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)

	stats := make(map[int]int)
	for _, n := range arr {
		if n > 0 && n%2 != 0 {
			stats[n]++
			continue
		}

		t := 0
		if n < 0 {
			t = n * 2
		} else {
			t = n / 2
		}

		if v, ok := stats[t]; ok {
			if v == 1 {
				delete(stats, t)
			} else {
				stats[t]--
			}
		} else {
			stats[n]++
		}
	}

	return len(stats) == 0
}
