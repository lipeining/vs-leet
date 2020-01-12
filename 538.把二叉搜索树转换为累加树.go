/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sum := 0
	// 反中序遍历
	var helper func(root *TreeNode)
	helper = func(root *TreeNode){
		if root == nil {
			return
		}
		helper(root.Right)
		sum += root.Val
		root.Val = sum
		helper(root.Left)
	}
	helper(root)
	return root
}
// @lc code=end

