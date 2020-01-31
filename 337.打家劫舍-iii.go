/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := root.Val
	if root.Left != nil {
		sum += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		sum += rob(root.Right.Left) + rob(root.Right.Right)
	}
	sum2 := rob(root.Left) + rob(root.Right)
	if sum2 > sum {
		return sum2
	}
	return sum
    // var dfs func(root *TreeNode) int
    // dfs = func(root *TreeNode) int {
	// 	if root == nil {
	// 		return 0
	// 	}
	// 	sum := root.Val
	// 	if root.Left != nil {
	// 		sum += dfs(root.Left.Left) + dfs(root.Left.Right)
	// 	}
	// 	if root.Right != nil {
	// 		sum += dfs(root.Right.Left) + dfs(root.Right.Right)
	// 	}
	// 	return sum
	// }
	// mid := dfs(root)
	// lr := 0
	// if root != nil {
	// 	lr = dfs(root.Left) + dfs(root.Right)
	// }
	// // [4,1,null,2,null,3]
	// // if root.Left != nil{
	// // 	lr += dfs(root.Left) 
	// // } 
	// // if root.Right != nil {
	// // 	lr += dfs(root.Right)
	// // }
	// if mid > lr {
	// 	return mid
	// }
	// return lr
}
// @lc code=end

