/*
 * @lc app=leetcode.cn id=872 lang=golang
 *
 * [872] 叶子相似的树
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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	l := inorder(root1)
	r := inorder(root2)
	if len(l) != len(r) {
		return false
	}
	for i := 0; i < len(l); i++ {
		if l[i] != r[i] {
			return false
		}
	}
	return true
}
func inorder(root *TreeNode) []int {
	ans := make([]int, 0)
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		if isLeaf(root) {
			ans = append(ans, root.Val)
		}
		helper(root.Right)
	}
	helper(root)
	return ans
}
func isLeaf(t *TreeNode) bool {
	return t.Left == nil && t.Right == nil
}

// @lc code=end

