/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, sum int) [][]int {
	ans := make([][]int, 0)
    var dfs func(root *TreeNode, path []int, target int)
    dfs = func(root *TreeNode, path []int, target int) {
		if root == nil {
			return
		}
		if isLeaf(root) {
			if root.Val == target {
				path = append(path, root.Val)
				ans = append(ans, path)
			}
		} else {
			t1 := make([]int, len(path))
			t2 := make([]int, len(path))
			copy(t1, path)
			copy(t2, path)
			t1 = append(t1, root.Val)
			t2 = append(t2, root.Val)
			dfs(root.Left, t1, target - root.Val)
			dfs(root.Right, t2, target - root.Val)
		}
	}
	path := make([]int, 0)
	dfs(root, path, sum)
	return ans
}
func isLeaf(t *TreeNode) bool {
	return t.Left == nil && t.Right == nil
}
// @lc code=end

