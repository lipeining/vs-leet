/*
 * @lc app=leetcode.cn id=965 lang=golang
 *
 * [965] 单值二叉树
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
func isUnivalTree(root *TreeNode) bool {
    if root == nil {
		return true
	}
	var dfs func(root *TreeNode, target int)  bool 
	dfs = func(root *TreeNode, target int) bool {
		if root == nil {
			return true
		}
		if root.Val != target {
			return false
		}
		return dfs(root.Left, target) && dfs(root.Right, target)
	}
	return dfs(root, root.Val)
}
// @lc code=end

