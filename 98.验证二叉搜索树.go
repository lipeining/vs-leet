/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	// 中序遍历水一下先
	prev := math.MaxInt64
	ans := true
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if prev == math.MaxInt64 {
			prev = root.Val
		} else if root.Val <= prev {
			ans = false
		}
		prev = root.Val
		inorder(root.Right)
	}
	inorder(root)
	return ans
	// ans := make([]int, 0)
	// var inorder func(root *TreeNode)
	// inorder = func(root *TreeNode) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	inorder(root.Left)
	// 	ans = append(ans, root.Val)
	// 	inorder(root.Right)
	// }
	// inorder(root)
	// fmt.Println(ans)
	// for i:=0;i<len(ans)-1;i++ {
	// 	if ans[i] >= ans[i+1] {
	// 		return false
	// 	}
	// }
	// return true
}
// @lc code=end

