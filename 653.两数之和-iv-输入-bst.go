/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
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
func findTarget(root *TreeNode, k int) bool {
	// if root == nil {
	// 	return false
	// }
	// if root.Val == k {
	// 	return true
	// }
	// return findTarget(root.Left, k) || findTarget(root.Right, k)
	// || findTarget(root.Left, k-root.Val)
	// || findTarget(root.Rigth, k-root.Right)
	m := make(map[int]bool)
	var dfs func(root *TreeNode, k int)bool
	dfs = func(root *TreeNode, k int)bool{
		if root == nil {
			return false
		}
		_,ok := m[k-root.Val] 
		if ok {
			return true
		}
		m[root.Val] = true
		return dfs(root.Left, k) || dfs(root.Right, k)
	}
	return  dfs(root, k)
}

// @lc code=end

