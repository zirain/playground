package pro112

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil &&
		root.Right == nil {
		return root.Val == sum
	}

	sum -= root.Val

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
