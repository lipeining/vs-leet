/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
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
func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode,0)
	queue = append(queue, root)
	for len(queue) != 0 {
		length:=len(queue)
		sum := 0
		for i:=0;i<length;i++ {
			tmp := queue[i]
			sum += tmp.Val
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
		queue = queue[length:]
		fmt.Println(sum, length)
		res = append(res, float64(sum) / float64(length))
	}
	return res
}
// @lc code=end

