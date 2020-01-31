/*
 * @lc app=leetcode.cn id=894 lang=golang
 *
 * [894] 所有可能的满二叉树
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
func allPossibleFBT(N int) []*TreeNode {
	memo := make(map[int][]*TreeNode)
	var helper func(N int) []*TreeNode
	helper = func(N int) []*TreeNode {
		if _,ok := memo[N]; !ok {
			list := make([]*TreeNode, 0)
			if N == 1 {
				t := &TreeNode{Val : 0}
				list = append(list, t)
			} else if N%2 == 1 {
				for i:=0;i<N;i++ {
					j := N-i-1
					for _,left := range helper(i) {
						for _,right := range helper(j) {
							t := &TreeNode{Val:0}
							t.Left = left
							t.Right = right
							list = append(list, t)
						}
					}				
				}
			}	
			memo[N] = list
		}
		return memo[N]
	}
	helper(N)
	return memo[N]
}
// @lc code=end

