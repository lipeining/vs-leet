/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 */
import "fmt"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}
	var dfs func(root *TreeNode, path string)
	dfs = func(root *TreeNode, path string) {
		str := strconv.Itoa(root.Val)
		p := str
		if path != "" {
			p = path + "->" + p
		}
		// fmt.Println(str, p, res)
		if root.Left == nil && root.Right == nil {
			res = append(res, p)
		} else if root.Left != nil && root.Right == nil {
			dfs(root.Left, p)
		} else if root.Left == nil && root.Right != nil {
			dfs(root.Right, p)
		} else {
			dfs(root.Left, p)
			dfs(root.Right, p)
		}
	}
	dfs(root, "")
	return res
}

// @lc code=end

