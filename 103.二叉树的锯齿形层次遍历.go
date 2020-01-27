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
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	left := true
	for len(queue) != 0 {
		length := len(queue)
		tmp := make([]int, 0)
		for i:=0;i<length;i++ {
			r := queue[i]
			if left {
				tmp = append(tmp, r.Val)
			} else {
				tmp = append([]int{r.Val}, tmp...)
			}
			if r.Left != nil {
				queue = append(queue, r.Left)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
			}
		}
		// fmt.Println(tmp)
		queue = queue[length:]
		left = !left
		ans = append(ans, tmp)
	}
	return ans
}
// @lc code=end

