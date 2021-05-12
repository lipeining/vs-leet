/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	if root == nil {
		return ans
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[i]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		ans = append(ans, level)
	}
	return ans
}

// @lc code=end

