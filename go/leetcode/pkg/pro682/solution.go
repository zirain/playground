package pro682

import "strconv"

func calPoints(ops []string) int {
	ans := 0

	stack := make([]int, 0, len(ops))
	for _, op := range ops {
		switch op {
		case "C":
			ans -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "D":
			prev := stack[len(stack)-1]
			stack = append(stack, prev*2)
			ans += prev * 2
		case "+":
			prev := stack[len(stack)-1]
			prevprev := stack[len(stack)-2]
			stack = append(stack, prev+prevprev)
			ans += prev + prevprev
		default:
			point, _ := strconv.ParseInt(op, 10, 32)
			stack = append(stack, int(point))
			ans += int(point)
		}
	}

	return ans
}
