/*
 * @lc app=leetcode.cn id=1302 lang=golang
 *
 * [1302] 层数最深叶子节点的和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	ans := 0
	maxDepth := 0
    var dfs func(root *TreeNode, depth int)
    dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		dfs(root.Left, depth + 1)
		dfs(root.Right, depth + 1)
		if depth > maxDepth {
			maxDepth = depth
			ans = root.Val
		} else if depth == maxDepth {
			ans += root.Val
		}
	}
	dfs(root, 0)
	return ans
}
// @lc code=end

