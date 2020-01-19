/*
 * @lc app=leetcode.cn id=897 lang=golang
 *
 * [897] 递增顺序查找树
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
func increasingBST(root *TreeNode) *TreeNode {
	ans := TreeNode{Val:-1}
	p := &ans
	var inorder func(root *TreeNode) 
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		tmp := TreeNode{Val: root.Val}
		p.Right = &tmp
		p = p.Right
		inorder(root.Right)
	}
	inorder(root)
	return ans.Right
}
// @lc code=end

