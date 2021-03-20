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
	sum := 0
	var dfs func(node *TreeNode, val int)
	dfs = func(node *TreeNode, val int) {
		if node == nil {
			return
		}
		next := val << 1 | node.Val
		if node.Left == nil && node.Right == nil {
			sum += next
		} else {
			dfs(node.Left, next)
			dfs(node.Right, next)
		}
	}
	dfs(root, 0)
	return sum % int(1e9+7)
}
func sumRootToLeaf2(root *TreeNode) int {
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

