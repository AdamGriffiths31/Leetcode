func sortedArrayToBST(nums []int) *TreeNode {
	return helper(0, len(nums)-1, nums)
}

func helper(left, right int, nums []int) *TreeNode {
	if left > right {
		return nil
	}

	middle := (left + right) / 2
	root := TreeNode{
		Val:   nums[middle],
		Left:  helper(left, middle-1, nums),
		Right: helper(middle+1, right, nums),
	}

	return &root
}