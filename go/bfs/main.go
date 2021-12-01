package main

import (
	"fmt"
	"os"
)

func readFromFile(filename string) [][]int{
	// 定义一个行和列,用来接收迷宫是几行几列的数组
	var row, col int
	file, e := os.Open(filename)
	if e != nil {
		panic("error")
	}
	defer file.Close()

	fmt.Fscan(file, &row, &col)

	// 定义一个数组
	// 注意: 定义数组的时候, 我们只要传入几行就可以了.
	// 二维数组的含义, 其实质是一个一维数组, 一维数组里每一个元素又是一个数组
	maze := make([][]int, row)
	for i := 0; i < len(maze); i++ {
		maze[i] = make([]int, col)
		for j := 0; j < len(maze[i]); j++ {
			fmt.Fscan(file, &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

// 当前节点, 向四个方向探索后的节点
// 这里使用的是返回新的节点的方式, 不修改原来的节点. 所以使用的是值拷贝,而不是传地址
func (p point) add(dir point) point{
	return point{p.i + dir.i, p.j + dir.j }
}

// 获取某个点的坐标值
// 同时判断这个点有没有越界, 返回的是这个值是否有效
// return 第一个参数表示返回的值是否是1, 是1表示撞墙了
//        第二个参数表示返回的值是否不越界, 不越界返回true, 越界,返回false 就和你
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j <0 || p.j >= len(grid[0]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

// 定义要探索的方向, 上下左右四个方向
var dirs = []point {
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

// 走迷宫
func walk(maze [][]int, start, end point) [][]int {
	// 第一步: step代表从start开始, 走了多少步走到目标点, 最后的路径是通过这个创建出来的, 最后从后往前推就可以算出最短路径
	// 2. 通step还可以知道哪些点是到过的, 哪些点是没到过的
	step := make([][]int, len(maze))
	for i := range step {
		// 定义每一行有多少列, 这样就定义了一个和迷宫一样的二维数组
		step[i] = make([]int, len(maze[i]))
	}

	// 第二步: 定义一个队列, 用来保存已经发现还未探索的点, 队列里的初始值是(0,0)点
	que := []point{start}

	// 第三步: 开始走迷宫, 走迷宫退出的条件有两个
	// 1. 走到终点, 退出
	// 2. 队列中没有元素, 退出
	for len(que) > 0 {
		// 开始探索, 依次取出队列中, 已经发现还未探索的元素
		// cur 表示当前要探索的节点
		cur := que[0]
		// 然后从头拿掉第一个元素
		que = que[1:]

		// 如果这个点是终点, 就不向下探索了
		if cur == end {
			break
		}

		// 当前节点怎么探索呢? 要往上下左右四个方向去探索
		for _, dir := range dirs {
			// 探索下一个节点, 这里获取下一个节点的坐标. 当前节点+方向
			next := cur.add(dir)
			// 判断坐标是否符合探索的要求
			// 1. maze at next is 0
			// 2. and step at next is 0, 如果step的值不是0 ,说明曾经到过这个点了, 不能重复走
			// 3. next != start 处理特殊点, (0,0)点

			// 探索这个点是否是墙
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			// 探索这个点是否已经走过
			val, ok = next.at(step)
			if val != 0 || !ok {
				continue
			}

			// 走到起始点了, 返回
			if next == start {
				continue
			}

			// 已经找到这个点了, 计算当前的步数
			curval, _ := cur.at(step) // 当前这一步的步数
			step[next.i][next.j] = curval + 1 // 下一步是当前步数+1
			que = append(que, next) // 将下一步节点放入到队列中
		}

	}
	return step
}

func main() {

	maze := readFromFile("maze/maze.in")

	// (len(maze)-1, len[maze[0]]-1) 是终点
	// (0, 0) 是起始点
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row  {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}