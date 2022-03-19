package pro606

import (
	"strconv"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func tree2str(root *TreeNode) string {
	ans := ""

	if root == nil {
		return "()"
	}

	ans += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return ans
	}
	ans += "("
	if root.Left != nil {
		ans += tree2str(root.Left)
	}
	ans += ")"

	if root.Right != nil {
		ans += "("
		ans += tree2str(root.Right)
		ans += ")"
	}

	return string(ans)
}
