/*
 * @lc app=leetcode.cn id=971 lang=golang
 *
 * [971] 翻转二叉树以匹配先序遍历
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
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	ans := make([]int, 0)
	index := 0
	flag := false
  var preorder func(r *TreeNode)
  preorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		// fmt.Println(r.Val)
		if r.Val != voyage[index] {
			flag = true
			return
		}
		index++
		if r.Left != nil {
			if r.Left.Val != voyage[index] {
				ans = append(ans, r.Val)
				preorder(r.Right)
				preorder(r.Left)
			} else {
				preorder(r.Left)
				preorder(r.Right)
			}
		} else {
			preorder(r.Left)
			preorder(r.Right)
		}
	}
	preorder(root)
	if flag {
		return []int{-1}
	}
	return ans
}
// @lc code=end

