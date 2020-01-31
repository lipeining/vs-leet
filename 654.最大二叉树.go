/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	index := 0
	max := nums[0]
	for i:=0;i<length;i++ {
		if nums[i] > max {
			max = nums[i]
			index=i
		}
	}
	root := &TreeNode{Val : max}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}
// @lc code=end

