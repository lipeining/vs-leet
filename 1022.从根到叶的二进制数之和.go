/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] 从根到叶的二进制数之和
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
func sumRootToLeaf(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode, path []int)
	dfs = func(root *TreeNode, path []int) {
		if root == nil {
			return
		}
		tmp := append(path, root.Val)
		if isLeaf(root) {
			ans += path2num(tmp)
		} else {
			dfs(root.Left, tmp)
			dfs(root.Right, tmp)
		}
	}
	path := make([]int, 0)
	dfs(root, path)
	return ans
}
func path2num(path []int) int {
	multi := 1
	n := 0
	fmt.Println(path)
	for i:=len(path)-1;i>=0;i-- {
		n += multi*path[i]
		multi = multi * 2
	}
	return n
}

func isLeaf(t *TreeNode)bool {
	return t.Left == nil && t.Right == nil
}
// @lc code=end

