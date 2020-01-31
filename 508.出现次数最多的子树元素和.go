/*
 * @lc app=leetcode.cn id=508 lang=golang
 *
 * [508] 出现次数最多的子树元素和
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
func findFrequentTreeSum(root *TreeNode) []int {
	counter := make(map[int]int)
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		sum := root.Val + dfs(root.Left) + dfs(root.Right)
		counter[sum]++
		return sum
	}
	dfs(root)

	ans := make([]int, 0)
	max := 0
	for _,v := range counter {
		if v > max {
			max = v
		}
	}
	for k,v := range counter {
		if v == max {
			ans = append(ans, k)
		}
	}
	fmt.Println(counter, max)
	return ans
}
// @lc code=end

