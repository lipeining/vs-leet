/*
 * @lc app=leetcode.cn id=979 lang=golang
 *
 * [979] 在二叉树中分配硬币
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
func distributeCoins(root *TreeNode) int {
	ans := 0
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l:=dfs(root.Left)
		r:=dfs(root.Right)
		ans += int(math.Abs(float64(l))) + int(math.Abs(float64(r)))
		return  root.Val + r + l -1
	}
	dfs(root)
	return ans
}

// @lc code=end

