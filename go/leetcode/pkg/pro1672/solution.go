package pro1672

func maximumWealth(accounts [][]int) int {
	max := 0

	for _, account := range accounts {
		wealth := sum(account)
		if wealth > max {
			max = wealth
		}
	}

	return max
}

func sum(nums []int) int {
	total := 0
	for _, t := range nums {
		total += t
	}
	return total
}
