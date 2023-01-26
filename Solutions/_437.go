type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int
var hashmap map[int]int

func pathSum(root *TreeNode, targetSum int) int {
	res = targetSum
	hashmap = make(map[int]int)
	hashmap[0] = 1

	return dfs(root, 0)

}

func dfs(root *TreeNode, rootSum int) int {
	if root == nil {
		return 0
	}
	rootSum += root.Val
	count := hashmap[rootSum-res]
	hashmap[rootSum]++

	count += dfs(root.Left, rootSum)
	count += dfs(root.Right, rootSum)
	hashmap[rootSum]--

	return count
}
