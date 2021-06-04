/*
 * @lc app=leetcode.cn id=802 lang=golang
 *
 * [802] 找到最终的安全状态
 */

// @lc code=start
func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	color := make([]int, n)
	ans := make([]int, 0)
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if color[i] > 0 {
			return color[i] == 2
		}
		color[i] = 1
		for _, nei := range graph[i] {
			if color[i] == 2 {
				continue
			}
			if color[nei] == 1 || !dfs(nei) {
				return false
			}
		}
		color[i] = 2
		return true
	}
	for i := 0; i < n; i++ {
		if dfs(i) {
			ans = append(ans, i)
		}
	}
	return ans
}

// @lc code=end

