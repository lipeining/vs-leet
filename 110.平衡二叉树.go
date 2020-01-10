/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 */
import (
	"fmt",
	"math"
)
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}
	leftDepth := depth(root.Left)
	rightDepth := depth(root.Right)
	diff := leftDepth - rightDepth
	return math.Abs(diff) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
func depth(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return math.Max(depth(root.Left) + 1, depth(root.Right) + 1)
}
// @lc code=end

