/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		size := len(ans)
		if size < depth+1 {
			// insert
			ans = append(ans, node.Val)
		} else {
			ans[depth] = node.Val
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	return ans
}
func rightSideViewBFS(root *TreeNode) []int {
	// 只看到最右侧的节点
	// bfs 最简单，取 length-1 即可，然后继续下去

	queue := make([]*TreeNode, 0)
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			t := queue[i]
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
		}
		ans = append(ans, queue[length-1].Val)
		queue = queue[length:]
		// fmt.Println(queue)
	}
	return ans
}

// @lc code=end

