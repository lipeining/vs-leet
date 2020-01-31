/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
func largestValues(root *TreeNode) []int {
	ans := make([]int,0)
	if root == nil {
		return ans
	}
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return 
		}
		length := len(ans)
		if length >= depth {
			if root.Val > ans[depth-1] {
				ans[depth-1] = root.Val
			}
		} else {
			ans = append(ans, root.Val)
		}
		dfs(root.Left, depth + 1)
		dfs(root.Right, depth + 1)
	}
	dfs(root, 1)
	return ans

	// // bfs 划水
	// queue := make([]*TreeNode, 0)
	// ans := make([]int, 0)
	// if root == nil {
	// 	return ans
	// }
	// queue = append(queue, root)
	// for len(queue) != 0 {
	// 	length := len(queue)
	// 	max := queue[0].Val
	// 	for i:=0;i<length;i++ {
	// 		t := queue[i]
	// 		if t.Val > max {
	// 			max = t.Val
	// 		}
	// 		if t.Left != nil {
	// 			queue = append(queue, t.Left)
	// 		}
	// 		if t.Right != nil {
	// 			queue = append(queue, t.Right)
	// 		}
	// 	}
	// 	ans = append(ans, max)
	// 	queue = queue[length:]
	// }
	// return ans
}
// @lc code=end

