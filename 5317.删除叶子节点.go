/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	// 从下而上的递归方法删除
	if root == nil {
		return root
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil  && root.Val == target{
		return nil
	}
	return root
}