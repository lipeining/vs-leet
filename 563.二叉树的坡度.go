/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
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
func findTilt(root *TreeNode) int {
	res:=float64(0)
	var dfs func(root *TreeNode)int
	dfs = func(root *TreeNode)int {
		if root == nil {
			return 0
		}
		
		l := dfs(root.Left)
		r := dfs(root.Right)
		res += math.Abs(float64(l-r))
		return l + r + root.Val
	}
	dfs(root)
	return int(res)
}
// @lc code=end

