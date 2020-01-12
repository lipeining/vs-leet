/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, sum int) int {
	res := 0
	var helper func(root *TreeNode, sum int)
    helper = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		tmp := sum - root.Val
		if tmp == 0 {
			res++
		}
		helper(root.Left, tmp)
		helper(root.Right, tmp)
	}
	helper(root, sum)
	return res
}
// @lc code=end

