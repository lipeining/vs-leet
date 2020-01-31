/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {

	counter:=0
	ans := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		counter++
		if counter == k {
			ans = root.Val
		}
		dfs(root.Right)
	}
	dfs(root)
	return ans


    // list := make([]int, 0)
	// var dfs func(root *TreeNode)
	// dfs = func(root *TreeNode) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	dfs(root.Left)
	// 	list = append(list, root.Val)
	// 	dfs(root.Right)
	// }
	// dfs(root)
	// return list[k-1]
}
// @lc code=end

