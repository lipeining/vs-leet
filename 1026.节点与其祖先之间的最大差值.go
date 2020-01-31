/*
 * @lc app=leetcode.cn id=1026 lang=golang
 *
 * [1026] 节点与其祖先之间的最大差值
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
func maxAncestorDiff(root *TreeNode) int {
	ans := -1
	if root == nil {
		return 0
	}
	var dfs func(root *TreeNode, max int, min int)
	dfs = func(root *TreeNode, max int, min int) {
		if root == nil {
			return
		}
		if root.Val > max {
			max = root.Val 
		}
		if root.Val < min {
			min = root.Val
		}
		if root.Left == nil && root.Right == nil {
			diff := int(math.Abs(float64(max - min)))
			if diff > ans {
				ans = diff
			}
		}
		dfs(root.Left, max, min)
		dfs(root.Right, max, min)
	}
	dfs(root, root.Val, root.Val)
	return ans
	// 超时了
	// ans := -1
	// if root == nil {
	// 	return 0
	// }
	// var dfs func(root *TreeNode, parent int)
	// dfs = func(root *TreeNode, parent int) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	diff := int(math.Abs(float64(parent - root.Val)))
	// 	if diff > ans {
	// 		ans = diff
	// 	}
	// 	dfs(root.Left, parent)
	// 	dfs(root.Right, parent)
	// 	dfs(root.Left, root.Val)
	// 	dfs(root.Right, root.Val)
	// }
	// dfs(root, root.Val)
	// return ans
}
// @lc code=end

