/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
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
func sumNumbers(root *TreeNode) int {
	ans := 0
	isLeaf := func(t *TreeNode) bool {
		return t.Left == nil && t.Right == nil
	}
	var dfs func(node *TreeNode, num int)
	dfs = func(node *TreeNode, num int) {
		if node == nil {
			return
		}
		num = num*10 + node.Val
		if isLeaf(node) {
			ans += num
			return
		}
		dfs(node.Left, num)
		dfs(node.Right, num)
	}
	dfs(root, 0)
	return ans
}

// @lc code=end

