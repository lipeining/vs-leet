/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, sum int) int {
	// [1,null,2,null,3,null,4,null,5]
	// 重复遍历了 3 这个节点，遍历了两次
	// res := 0
	// var helper func(root *TreeNode, currenct int, total int)
    // helper = func(root *TreeNode, currenct int, total int) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	tmp := currenct - root.Val
	// 	fmt.Println(root, currenct, root.Val, total)
	// 	if tmp == 0 {
	// 		res++
	// 	}
	// 	helper(root.Left, tmp, total)
	// 	helper(root.Right, tmp, total)
	// 	helper(root.Left, total, total)
	// 	helper(root.Right, total, total)
	// }
	// helper(root, sum, sum)
	// return res
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode, sum int) int
	dfs = func(node *TreeNode, sum int) int {
		if node == nil {
			return 0
		}
		count := 0
		if sum == node.Val {
			count = 1
		}
		return dfs(node.Left, sum - node.Val) + dfs(node.Right, sum - node.Val) + count
	}
	return dfs(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
// @lc code=end

