package pro589

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		ans = append(ans, cur.Val)
		for i := len(cur.Children) - 1; i >= 0; i-- {
			queue = append(queue, cur.Children[i])
		}
	}

	return ans
}
