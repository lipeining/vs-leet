/*
 * @lc app=leetcode.cn id=951 lang=golang
 *
 * [951] 翻转等价二叉树
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
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
		return true
	} else if root1 != nil && root2 == nil {
		return false
	} else if root1 == nil && root2 != nil {
		return false
	} else {
		if root1.Val != root2.Val {
			return false
		}
		return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) || flipEquiv(root1.Right, root2.Left) && flipEquiv(root1.Left, root2.Right))
	}
}
// @lc code=end

