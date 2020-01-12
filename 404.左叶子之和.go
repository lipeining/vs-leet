/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
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
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
		return 0
	}
	left := root.Left
	res := 0
	if left != nil && left.Left == nil && left.Right == nil{
		res += left.Val
	} else {
		res += sumOfLeftLeaves(root.Left)
	}
	res += sumOfLeftLeaves(root.Right)
	return res
	
}
// @lc code=end

