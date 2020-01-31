/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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
func generateTrees(n int) []*TreeNode {
	// 官方思路
	if n == 0 {
		ans := make([]*TreeNode, 0)
		return ans
	}
	return helper(1, n)
}
func helper(start int, end int) []*TreeNode {
	trees := make([]*TreeNode, 0)
	if start > end {
		// trees = make([]*TreeNode, 1)
		trees = append(trees , nil)
		return trees
	}
	for i:=start;i<=end;i++ {
		lefts := helper(start, i-1)
		rights := helper(i+1, end)
		for _,l := range lefts {
			for _,r := range rights {
				// cur := &TreeNode{i, nil, nil}
				cur := &TreeNode{Val : i}
				cur.Left = l
				cur.Right = r
				trees = append(trees, cur)
			}
		}
	}
	return trees
}
// @lc code=end

