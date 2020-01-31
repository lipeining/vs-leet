/*
 * @lc app=leetcode.cn id=1008 lang=golang
 *
 * [1008] 先序遍历构造二叉树
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
func bstFromPreorder(preorder []int) *TreeNode {
	length := len(preorder)
	if length == 0 {
		return nil
	}
	// if length == 1 {
	// 	t := &TreeNode{Val :preorder[0]}
	// 	return t
	// }
	head := preorder[0]
	index := 0
	for i:=0;i<length;i++ {
		if preorder[i] > head {
			index = i
			break
		}
	}
	root := &TreeNode{Val:head}
	if index >= 1 {
		root.Left = bstFromPreorder(preorder[1:index])
		root.Right = bstFromPreorder(preorder[index:])
	} else {
		root.Left = bstFromPreorder(preorder[1:])
	}
	return root
}
// @lc code=end

