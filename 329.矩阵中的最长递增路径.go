/*
 * @lc app=leetcode.cn id=329 lang=golang
 *
 * [329] 矩阵中的最长递增路径
 */

// @lc code=start
func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		// memo[i][j]++
		now := 1
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && matrix[ni][nj] > matrix[i][j] {
				now = max(now, dfs(ni, nj)+1)
			}
		}
		memo[i][j] = now
		return memo[i][j]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	fmt.Println("anssss", ans)
	return ans
}

// @lc code=end

