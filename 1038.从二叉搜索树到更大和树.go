/*
 * @lc app=leetcode.cn id=1038 lang=golang
 *
 * [1038] 从二叉搜索树到更大和树
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
func bstToGst(root *TreeNode) *TreeNode {
	now := 0
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Right)
		now += root.Val
		root.Val = now
		inorder(root.Left)
	}
	inorder(root)
	return root
	// list := make([]int, 0)
	// var inorder func(root *TreeNode)
	// inorder = func(root *TreeNode) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	inorder(root.Left)
	// 	list = append(list, root.Val)
	// 	inorder(root.Right)
	// }
	// inorder(root)
	// // 处理list
	// for i := len(list) - 2; i >= 0; i-- {
	// 	list[i] += list[i+1]
	// }
	// // fmt.Println(list)
	// cur := 0
	// var helper func(root *TreeNode)
	// helper = func(root *TreeNode) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	helper(root.Left)
	// 	// list 赋值
	// 	root.Val = list[cur]
	// 	cur++
	// 	helper(root.Right)
	// }
	// helper(root)
	// return root
}

// @lc code=end
