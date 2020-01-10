/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
 */
import "fmt"
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	// append 方法构建一个倒序的层次遍历
	res := make([][]int, 0)
	// queue := []*TreeNode{
	// 	root
	// }
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
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
		// fmt.Println("level", level)
		// res = append(res, level)
		res = append([][]int{level}, res...)
	}
	// fmt.Println(res)
	return res
	// result := make([][]int, len(res))
	// for i:=len(res)-1;i>=0;i-- {
	// 	result[len(res)-1-i] = res[i]
	// }
	// return result
}
// @lc code=end

