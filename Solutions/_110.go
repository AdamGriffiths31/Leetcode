func isBalanced(root *TreeNode) bool {
	result, _ := dfs(root)
	return result
}

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	leftBalanced, leftHeight := dfs(root.Left)
	rightBalanced, rightHeight := dfs(root.Right)

	balance := (rightBalanced && leftBalanced) && (abs(leftHeight-rightHeight) <= 1)

	return balance, max(leftHeight, rightHeight) + 1
}

func max(left int, right int) int {
	if left > right {
		return left
	}

	return right
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}