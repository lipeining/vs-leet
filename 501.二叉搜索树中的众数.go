/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
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
func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	smap := make(map[int]int)
	max := 1
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		c,ok:= smap[root.Val]
		if ok {
			next := c + 1
			smap[root.Val] = next
			if max < next {
				max = next
			}
		} else {
			smap[root.Val] = 1
		}
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)
	for k,v := range smap {
		if v==max {
			res = append(res, k)
		}
	}
	return res
}
// @lc code=end

