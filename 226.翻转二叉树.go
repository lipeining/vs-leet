/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return root
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	// 下面这两步需要通过画图得到
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
// @lc code=end

