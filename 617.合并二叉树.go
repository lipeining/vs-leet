/*
 * @lc app=leetcode.cn id=617 lang=golang
 *
 * [617] 合并二叉树
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
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 != nil && t2 == nil {
		return t1
	} else if t1 == nil && t2 != nil {
		return t2
	} else {
		// v := t1.Val + t2.Val
		// tmp := TreeNode{v, nil, nil}
		tmp :=  TreeNode{Val : t1.Val+t2.Val}
		tmp.Left = mergeTrees(t1.Left, t2.Left)
		tmp.Right = mergeTrees(t1.Right, t2.Right)
		return &tmp
	}
}
// @lc code=end

