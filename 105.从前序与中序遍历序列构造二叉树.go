/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 知道的思路是 前序中找到中序中的左子树分支和右子树的划分点
	length := len(preorder)
	var root *TreeNode
	if length == 0 {
		return root
	}
	head := &TreeNode{Val: preorder[0]}
	mid := 0
	for i:=0;i<len(inorder);i++ {
		if inorder[i] == preorder[0] {
			mid = i
			break
		}
	}
	head.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	head.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return head
}
// @lc code=end

