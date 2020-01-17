/*
 * @lc app=leetcode.cn id=671 lang=golang
 *
 * [671] 二叉树中第二小的节点
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
func findSecondMinimumValue(root *TreeNode) int {
	// var helper func(root *TreeNode, min int) int
	// helper = func(root *TreeNode, min int) int {
	// 	if root == nil {
	// 		return -1
	// 	}
	// 	if root.Left!=nil || root.Right!=nil {
	// 		min := root.Left.Val
	// 		min2 := root.Right.Val
	// 		if min < root.Right.Val {
	// 			min = root.Right.Val
	// 			min2 = root.Left.Val
	// 		}
	// 		if root.Val == min {
	// 			if root.Left.Val == root.Right.Val {
	// 				// 三个数都相等
	// 				l := helper(root.Left, min)
	// 				r := helper(root.Right, min)
	// 				if l <= r {
	// 					return l
	// 				}
	// 			} else {
	// 				return min2
	// 			} 
	// 		}
	// 		return min
	// 	} else {
	// 		return -1
	// 	}
	// }
	// return helper(root, 0)
	// 用闭包比较好
	if root == nil {
		return -1
	}
	res := root.Val
	base := root.Val
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root != nil {
			if root.Val > base {
				if base == res {
					res = root.Val
				} else if root.Val < res {
					res = root.Val
				}
				return;
			}
			helper(root.Left)
			helper(root.Right)
		}
	}
	helper(root)
	if res == base {
		return -1
	}
	return res
}
// @lc code=end

