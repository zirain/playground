package pro429

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)

	q := []*Node{root}

	for len(q) > 0 {
		level := q
		q = nil
		tmp := make([]int, 0, len(level))
		for _, node := range level {
			if len(node.Children) != 0 {
				q = append(q, node.Children...)
			}
			tmp = append(tmp, node.Val)
		}
		ans = append(ans, tmp)
	}

	return ans
}
