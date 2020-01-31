/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
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
func findBottomLeftValue(root *TreeNode) int {
	maxDepth := 0
	ans := 0
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		dfs(root.Left, depth + 1)
		if depth > maxDepth {
			maxDepth = depth
			ans = root.Val
		}
		dfs(root.Right, depth + 1)
	}
	dfs(root, 1)
	return ans


	// // bfs 划水
	// queue := make([]*TreeNode, 0)
	// if root == nil {
	// 	return 0
	// }
	// queue = append(queue, root)
	// ans := 0
	// for len(queue) != 0 {
	// 	length := len(queue)
	// 	for i:=0;i<length;i++ {
	// 		t := queue[i]
	// 		if i == 0 {
	// 			ans = t.Val
	// 		}
	// 		if t.Left != nil {
	// 			queue = append(queue, t.Left)
	// 		}
	// 		if t.Right != nil {
	// 			queue = append(queue, t.Right)
	// 		}
	// 	}
	// 	queue = queue[length:]
	// }
	// return ans
}
// @lc code=end

