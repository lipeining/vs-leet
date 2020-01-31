/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
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
func sumNumbers(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum = sum * 10 + root.Val
		if root.Left == nil && root.Right == nil {
			ans+=sum
			return
		} else {
			dfs(root.Left, sum)
			dfs(root.Right, sum)
		}
	}
	dfs(root, 0)
	return ans
}
// @lc code=end

