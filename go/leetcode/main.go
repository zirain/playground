package main

import "fmt"

func get(x int) int {
	res := 0
	for x > 0 {
		res += x % 10
		x = x / 10
	}
	return res
}

func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}
	vis := make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]int, n)
	}
	ans := 1
	vis[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 && j == 0) || get(i)+get(j) > k {
				continue
			}

			// 边界判断
			if i-1 >= 0 {
				vis[i][j] |= vis[i-1][j]
			}
			if j-1 >= 0 {
				vis[i][j] |= vis[i][j-1]
			}
			ans += vis[i][j]
		}
	}
	return ans
}

func main() {
	fmt.Println(movingCount(2, 3, 1))
	fmt.Println(movingCount(2, 3, 0))
}
