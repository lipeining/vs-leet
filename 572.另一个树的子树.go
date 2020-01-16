/*
 * @lc app=leetcode.cn id=572 lang=golang
 *
 * [572] 另一个树的子树
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
func isSubtree(s *TreeNode, t *TreeNode) bool {
	var helper func(s *TreeNode, t *TreeNode)bool
	helper = func(s *TreeNode, t *TreeNode)bool {
		if s == nil && t == nil {
			return true
		} else if s == nil && t != nil {
			return false
		} else if s != nil && t == nil {
			return false
		} else if s.Val != t.Val {
			return false
		} else {
			return helper(s.Left, t.Left) && helper(s.Right, t.Right)
		}
	}
	if s == nil {
		return false
	}
	if t == nil {
		return true
	}
	if s.Val != t.Val {
		return isSubtree(s.Left, t) || isSubtree(s.Right, t)
	}
	return helper(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
	// if s == nil && t == nil {
	// 	return true
	// } else if s == nil && t != nil {
	// 	return false
	// } else if s != nil && t == nil {
	// 	return false
	// } else if s.Val != t.Val {
	// 	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
	// } else {
	// 	return isSubtree(s.Left, t.Left) && isSubtree(s.Right, t.Right)
	// }
	// return isSubtree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}
// @lc code=end

