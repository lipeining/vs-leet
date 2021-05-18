/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := max(maxGain(node.Left), 0)
		r := max(maxGain(node.Right), 0)
		p := node.Val + l + r
		ans = max(ans, p)
		return node.Val + max(l, r)
	}
	maxGain(root)
	return ans
}

// @lc code=end

