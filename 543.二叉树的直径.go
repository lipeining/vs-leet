/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		L := depth(root.Left)
		R := depth(root.Right)
		tmp := L + R + 1
		if tmp > ans {
			ans = tmp
		}
		return 1 + max(L, R)
	}
	depth(root)
	return ans - 1
}

// @lc code=end

