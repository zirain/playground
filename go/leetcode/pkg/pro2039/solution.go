package pro2039

func networkBecomesIdle(edges [][]int, patience []int) (ans int) {
	n := len(patience)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	vis := make([]bool, n)
	vis[0] = true
	q := []int{0}
	for dist := 1; q != nil; dist++ {
		tmp := q
		q = nil
		for _, x := range tmp {
			for _, v := range g[x] {
				if vis[v] {
					continue
				}
				vis[v] = true
				q = append(q, v)
				ans = max(ans, (dist*2-1)/patience[v]*patience[v]+dist*2+1)
			}
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
