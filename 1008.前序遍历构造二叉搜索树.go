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
	n := len(preorder)
	if n == 0 {
		return nil
	}
	head := preorder[0]
	root := &TreeNode{
		Val: head,
	}
	i := 1
	for i < n {
		if preorder[i] > head {
			break
		}
		i++
	}
	root.Left = bstFromPreorder(preorder[1:i])
	root.Right = bstFromPreorder(preorder[i:])
	return root
}

// @lc code=end

