/*
 * @lc app=leetcode.cn id=1123 lang=golang
 *
 * [1123] 最深叶节点的最近公共祖先
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
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	max := func(a, b int)int{
		if a > b {
			return a
		}
		return b
	}
	var dfs func(node *TreeNode)int
	dfs = func(node *TreeNode)int{
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		return max(l, r) + 1
	}
	if root == nil {
		return root
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	if l == r {
		return root
	} else if l < r {
		return lcaDeepestLeaves(root.Right)
	} else {
		return lcaDeepestLeaves(root.Left)
	}
}
func lcaDeepestLeavesN(root *TreeNode) *TreeNode {
	_, ans := helper(root)
	return ans
}
func helper(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}
	l, lp := helper(root.Left)
	r, rp := helper(root.Right)
	if l == r {
		return l + 1, root
	} else if l < r {
		return r + 1, rp
	}
	return l + 1, lp
}

// @lc code=end

