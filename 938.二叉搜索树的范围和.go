/*
 * @lc app=leetcode.cn id=938 lang=golang
 *
 * [938] 二叉搜索树的范围和
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
func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val < L {
			dfs(root.Right)
		} else if root.Val > R {
			dfs(root.Left)
		} else {
			sum += root.Val
			dfs(root.Left)
			dfs(root.Right)
		}
	}
	dfs(root)
	return sum
}
// @lc code=end

