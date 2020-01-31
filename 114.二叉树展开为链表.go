/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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

func flatten(root *TreeNode) {
	var pre *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		dfs(root.Left)
		root.Right = pre
		root.Left = nil
		pre = root
	}
	dfs(root)
}
// func flatten(root *TreeNode)  {
//     if root == nil {
// 		return
// 	}
// 	root = dfs(root)
// }
// func dfs(root *TreeNode) *TreeNode {
// 	if root.Left == nil && root.Right == nil {
// 		return root
// 	} else if root.Left != nil && root.Right == nil {
// 		root.Right = dfs(root.Left)
// 		root.Left = nil
// 		return root
// 	} else if root.Left == nil && root.Right != nil {
// 		root.Right = dfs(root.Right)
// 		return root
// 	} else {
// 		left := dfs(root.Left)
// 		right := dfs(root.Right)
// 		// left 后面如何连接到 right
// 		root.Left = nil
// 		root.Right = left
// 		left.Right = right
// 		return root
// 	}
// }
// func isLeaf(t *TreeNode) bool {
// 	return t.Left == nil && t.Right == nil
// }
// @lc code=end

