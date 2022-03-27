package pro653

func findTarget(root *TreeNode, k int) bool {
	arr := []int{}
	var inorderTraversal func(*TreeNode)
	inorderTraversal = func(node *TreeNode) {
		if node != nil {
			inorderTraversal(node.Left)
			arr = append(arr, node.Val)
			inorderTraversal(node.Right)
		}
	}
	inorderTraversal(root)

	left, right := 0, len(arr)-1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == k {
			return true
		}
		if sum < k {
			left++
		} else {
			right--
		}
	}
	return false
}
