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
	max := 1
	var depth func(root *TreeNode)int 
	depth = func(root *TreeNode)int {
		if root == nil {
			return 0
		}
		L := depth(root.Left)
		R := depth(root.Right)
		tmp := L+R+1
		if tmp > max {
			max = tmp
		}
		return 1 + int(math.Max(float64(L),float64(R)))
	}
	depth(root)
	return max-1
}

// @lc code=end

