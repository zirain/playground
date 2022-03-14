package pro599

func findRestaurant(list1 []string, list2 []string) []string {
	minIndex := len(list1) + len(list2) + 1
	ans := make([]string, 0)
	mp := make(map[string]int)
	for i, s := range list1 {
		mp[s] = i
	}

	for i, s := range list2 {
		if idx, ok := mp[s]; ok {
			total := idx + i
			if total < minIndex {
				minIndex = total
				ans = []string{s}
			} else if total == minIndex {
				ans = append(ans, s)
			}

		}
	}

	return ans
}
