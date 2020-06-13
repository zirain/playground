package climbingstairs

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	prev, prevprev := 1, 1

	ans := 0
	for i := 2; i <= n; i++ {
		ans = prev + prevprev
		prevprev = prev
		prev = ans
	}
	return ans
}
