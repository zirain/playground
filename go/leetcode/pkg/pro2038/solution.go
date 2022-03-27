package pro2038

func winnerOfGame(colors string) bool {
	freq := [2]int{}
	cur, cnt := 'C', 0
	for _, c := range colors {
		if c != cur {
			cur, cnt = c, 1
		} else {
			cnt++
			if cnt >= 3 {
				freq[cur-'A']++
			}
		}
	}
	return freq[0] > freq[1]
}
