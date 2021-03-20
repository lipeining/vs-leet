/*
 * @lc app=leetcode.cn id=1038 lang=golang
 *
 * [1038] 把二叉搜索树转换为累加树
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
func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sum := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode){
		if node == nil {
			return
		}
		dfs(node.Right)
		node.Val += sum
		sum = node.Val
		dfs(node.Left)
	}
	dfs(root)
	return root
}
// @lc code=end

