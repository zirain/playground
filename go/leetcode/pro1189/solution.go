package pro1189

func maxNumberOfBalloons(text string) int {
	chrStats := []int{0, 0, 0, 0, 0}
	for _, chr := range text {
		switch chr {
		case 'b':
			chrStats[0]++
		case 'a':
			chrStats[1]++
		case 'l':
			chrStats[2]++
		case 'o':
			chrStats[3]++
		case 'n':
			chrStats[4]++
		}

	}

	return countBallon(chrStats)
}

func countBallon(stats []int) int {
	tar := []int{1, 1, 2, 2, 1}
	ans := 0
	for {
		for i := 0; i < len(tar); i++ {
			stats[i] -= tar[i]
			if stats[i] < 0 {
				return ans
			}
		}
		ans++
	}
}
