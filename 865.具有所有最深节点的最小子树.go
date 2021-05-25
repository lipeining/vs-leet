/*
 * @lc app=leetcode.cn id=865 lang=golang
 *
 * [865] 具有所有最深结点的最小子树
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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var ans *TreeNode
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	parent := make(map[*TreeNode]*TreeNode)
	set := make([]*TreeNode, 0)
	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := getDepth(node.Left)
		r := getDepth(node.Right)
		return 1 + max(l, r)
	}
	depth := getDepth(root)
	var dfs func(node, p *TreeNode, d int)
	dfs = func(node, p *TreeNode, d int) {
		if node == nil {
			return
		}
		parent[node] = p
		if d == depth {
			set = append(set, node)
		}
		dfs(node.Left, node, d+1)
		dfs(node.Right, node, d+1)
	}
	dfs(root, nil, 1)
	// fmt.Println("全部的最深节点", set)
	seen := make(map[*TreeNode]bool)
	for len(set) > 1 {
		size := len(set)
		for i := 0; i < size; i++ {
			node := set[i]
			p := parent[node]
			if !seen[p] {
				seen[p] = true
				set = append(set, p)
			}
		}
		set = set[size:]
	}
	ans = set[0]
	return ans
}

// @lc code=end

