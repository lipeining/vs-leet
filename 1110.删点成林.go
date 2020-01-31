/*
 * @lc app=leetcode.cn id=1110 lang=golang
 *
 * [1110] 删点成林
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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	ans := make([]*TreeNode, 0)
	if root == nil {
		return ans
	}	
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		root.Left = dfs(root.Left)
		root.Right = dfs(root.Right)
		if indexOf(to_delete, root.Val) {
			if root.Left != nil {
				ans = append(ans, root.Left)
			}
			if root.Right != nil {
				ans = append(ans, root.Right)
			}
			return nil
		} else {
			return root
		}
	}
	head := dfs(root)
	if head != nil {
		ans = append(ans, head)
	}
	// 由底向上进行删除操作：即返回 nil
	return ans
}
func indexOf(to_delete []int, target int) bool {
	for i:=0;i<len(to_delete);i++ {
		if to_delete[i] == target {
			return true
		}
	}
	return false
}
// @lc code=end

