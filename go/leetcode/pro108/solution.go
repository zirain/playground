package pro108

func sortedArrayToBST(nums []int) *TreeNode {
	c := len(nums)
	if c == 0 {
		return nil
	}

	mid := c / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
	return root
}
