var res int

func diameterOfBinaryTree(root *TreeNode) int {
    res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) int {
	if root == nil {
		return -1
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	res = max(res, 2+left+right)

	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}