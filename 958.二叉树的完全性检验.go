/*
 * @lc app=leetcode.cn id=958 lang=golang
 *
 * [958] 二叉树的完全性检验
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
func isCompleteTree(root *TreeNode) bool {
	// bfs 划水吧
	// 先得到 depth 
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	leaf := false
	for len(queue) != 0 {
		t := queue[0]
		queue = queue[1:]
		if t.Left == nil && t.Right != nil {
			return false
		}
		if leaf && (t.Left!=nil || t.Right !=nil) {
			return false
		}
		if t.Left !=nil {
			queue = append(queue, t.Left)
		}
		if t.Right !=nil {
			queue = append(queue, t.Right)
		} else {
			leaf = true
		}
	}
	return true
}
func depth(root *TreeNode)int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if left > right {
		return 1 + left
	}
	return 1 + right
}
// @lc code=end

