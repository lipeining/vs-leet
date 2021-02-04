/*
 * @lc app=leetcode.cn id=968 lang=golang
 *
 * [968] 监控二叉树
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
func minCameraCover(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 1
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left == 1 && right == 1 {
			return 0
		}
		if left+right >= 3 {
			return 1
		}
		ans++
		return 2
	}
	if dfs(root) == 0 {
		ans++
	}
	return ans
}

// @lc code=end

