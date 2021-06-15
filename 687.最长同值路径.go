/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
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
func longestUnivaluePath(root *TreeNode) int {
	ans := 0
	var arrowDfs func(root *TreeNode) int
	arrowDfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftLength := arrowDfs(root.Left)
		rightLength := arrowDfs(root.Right)
		l, r := 0, 0
		if root.Left != nil && root.Left.Val == root.Val {
			l = leftLength + 1
		}
		if root.Right != nil && root.Right.Val == root.Val {
			r = rightLength + 1
		}
		sum := l + r
		if sum > ans {
			ans = sum
		}
		if l > r {
			return l
		}
		return r
	}
	arrowDfs(root)
	return ans
}

// @lc code=end

