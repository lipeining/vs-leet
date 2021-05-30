/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	ans := make([][]int, 0)
	n := len(graph)
	var dfs func(index int, path []int)
	dfs = func(index int, path []int) {
		if index == n-1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for _, next := range graph[index] {
			path = append(path, next)
			dfs(next, path)
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	path = append(path, 0)
	dfs(0, path)
	return ans
}

// @lc code=end

