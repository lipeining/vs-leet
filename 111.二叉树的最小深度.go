/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 */
import "math"
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	// leftMin := min(root.Left)
	// rightMin := min(root.Right)
	// if int(leftMin) == 0  && rightMin > 0 {
	// 	return int(rightMin) + 1
	// } else if leftMin > 0 && int(rightMin) == 0 {
	// 	return int(leftMin) + 1
	// } else if int(leftMin) == 0 && int(rightMin) == 0 {
	// 	return 1
	// } else {
	// 	return int(math.Min(leftMin, rightMin)) + 1
	// }
	if root.Left == nil  && root.Right == nil {
		return 1
	}
	m := float64(1<<31-1)
	if root.Left != nil {
		m = math.Min(m, float64(minDepth(root.Left)))
	}
	if root.Right != nil {
		m = math.Min(m, float64(minDepth(root.Right)))
	}
	return int(m) + 1
}
// func min(root *TreeNode) float64 {
// 	if root == nil {
// 		return 0
// 	}
// 	return math.Min(
// 		float64(minDepth(root.Left)+1),
// 		float64(minDepth(root.Right)+1),
// 	)
// }
// @lc code=end

