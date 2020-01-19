/*
 * @lc app=leetcode.cn id=993 lang=golang
 *
 * [993] 二叉树的堂兄弟节点
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
func isCousins(root *TreeNode, x int, y int) bool {
	// 层次遍历法子吧
	if root == nil {
		return false
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for len(queue) != 0 {
		length := len(queue)
		counter := 0
		for i:=0; i<length;i++ {
			tmp := queue[i]
			if tmp.Val == x || tmp.Val == y {
				counter++
			}
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
			if tmp.Left != nil && tmp.Right != nil {
				if tmp.Left.Val == x && tmp.Right.Val == y {
					return false
				}
				if tmp.Left.Val == y && tmp.Right.Val == x {
					return false
				}
			}
		}
		if counter == 2 {
			return true
		}
		if counter == 1 {
			return false
		}
		queue = queue[length:]
	}
	return false
}
// @lc code=end

