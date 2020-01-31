/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	length := len(postorder)
	var root *TreeNode
	if length == 0 {
		return root
	}
	head := &TreeNode{Val: postorder[length-1]}
	mid := 0
	for i:=0;i<len(inorder);i++ {
		if inorder[i] == postorder[length-1] {
			mid = i
			break
		}
	}
	head.Left = buildTree(inorder[:mid], postorder[:mid])
	head.Right = buildTree(inorder[mid+1:], postorder[mid:length-1])
	return head
}
// @lc code=end

