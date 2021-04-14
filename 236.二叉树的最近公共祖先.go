/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
func lowestCommonAncestorN(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode
	var dfs func(root, p, q *TreeNode) int
	dfs = func(root, p, q *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left, p, q)
		right := dfs(root.Right, p, q)
		mid := 0
		if root == p || root == q {
			mid = 1
		}
		if left+right+mid >= 2 {
			ans = root
		}
		if left+right+mid > 0 {
			return 1
		}
		return 0
	}
	dfs(root, p, q)
	return ans
}

// @lc code=end

