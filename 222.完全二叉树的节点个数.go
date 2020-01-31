/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
	// ans := 0
	// var dfs func(root *TreeNode)
	// dfs = func(root *TreeNode) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	dfs(root.Left)
	// 	ans += 1
	// 	dfs(root.Right)
	// }
	// dfs(root)
	// return ans
	// queue := make([]*TreeNode, 0)
	// ans := 0
	// if root == nil {
	// 	return ans
	// }
	// queue = append(queue, root)
	// for len(queue) != 0 {
	// 	length := len(queue)
	// 	for i:=0;i<length;i++ {
	// 		t := queue[i]
	// 		if t.Left != nil {
	// 			queue = append(queue, t.Left)
	// 		}
	// 		if t.Right != nil {
	// 			queue = append(queue, t.Right)
	// 		}
	// 	}
	// 	ans+=length
	// 	queue = queue[length:]
	// 	// fmt.Println(queue)
	// }
	// return ans
}
// @lc code=end

