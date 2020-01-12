/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {
	// 只想到遍历的方法
	var inorder func(root *TreeNode)
	res := make([]int, 0)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	min := 1 << 31
	for i:=0; i < len(res)-1;i++ {
		diff := int(math.Abs(float64(res[i+1]-res[i])))
		if min > diff {
			min = diff
		}
	}
	return min
}
// @lc code=end

