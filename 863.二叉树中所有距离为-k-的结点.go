/*
 * @lc app=leetcode.cn id=863 lang=golang
 *
 * [863] 二叉树中所有距离为 K 的结点
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
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	ans := make([]int, 0)
	parent := make(map[*TreeNode]*TreeNode)
	var dfs func(node, p *TreeNode)
	dfs = func(node, p *TreeNode) {
		if node == nil {
			return
		}
		parent[node] = p
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)
	queue := make([]*TreeNode, 0)
	queue = append(queue, target)
	dis := 0
	seen := make(map[*TreeNode]bool)
	seen[target] = true
	for len(queue) != 0 {
		size := len(queue)
		if dis == K {
			for i := 0; i < size; i++ {
				ans = append(ans, queue[i].Val)
			}
		} else {
			for i := 0; i < size; i++ {
				t := queue[i]
				if t.Left != nil {
					if !seen[t.Left] {
						queue = append(queue, t.Left)
						seen[t.Left] = true
					}
				}
				if t.Right != nil {
					if !seen[t.Right] {
						queue = append(queue, t.Right)
						seen[t.Right] = true
					}
				}
				p := parent[t]
				if p != nil {
					if !seen[p] {
						queue = append(queue, p)
						seen[p] = true
					}
				}
			}
		}
		dis++
		queue = queue[size:]
	}
	return ans
}

// @lc code=end

