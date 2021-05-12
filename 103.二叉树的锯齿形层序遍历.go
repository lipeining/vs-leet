/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	if root == nil {
		return ans
	}
	reverse := func(num []int) []int {
		j := len(num) - 1
		i := 0
		for i < len(num)/2 {
			num[i], num[j] = num[j], num[i]
			i++
			j--
		}
		return num
	}
	queue = append(queue, root)
	left2right := true
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
		if left2right {
			ans = append(ans, level)
		} else {
			ans = append(ans, reverse(level))
		}
		left2right = !left2right
	}
	return ans
}

// @lc code=end

