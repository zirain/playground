package pro1447

import "fmt"

func simplifiedFractions(n int) []string {
	res := make([]string, 0)
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcd(i, j) != 1 {
				continue
			}

			res = append(res, fmt.Sprintf("%d/%d", i, j))
		}
	}

	return res
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}

	return gcd(b, a%b)
}
