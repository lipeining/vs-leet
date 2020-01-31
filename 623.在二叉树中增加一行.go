/*
 * @lc app=leetcode.cn id=623 lang=golang
 *
 * [623] 在二叉树中增加一行
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
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth == d-1 {
			oriLeft := root.Left
			tmpLeft := &TreeNode{Val : v}
			tmpLeft.Left = oriLeft
			
			oriRight := root.Right
			tmpRight := &TreeNode{Val : v}
			tmpRight.Right = oriRight

			root.Left = tmpLeft
			root.Right = tmpRight
			return
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	if d > 1 {
		dfs(root, 1)
	} else {
		tmpHead := &TreeNode{Val: v}
		tmpHead.Left = root
		return tmpHead
	}
	return root
	// bfs 可以划水吗？
}
// @lc code=end

