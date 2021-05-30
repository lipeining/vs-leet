/*
 * @lc app=leetcode.cn id=1034 lang=golang
 *
 * [1034] 边框着色
 */

// @lc code=start
func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	source := grid[r0][c0]
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 中心不需要渲染，边界需要渲染
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		if visited[i][j] {
			return 1
		}
		if grid[i][j] != source {
			return 0
		}
		visited[i][j] = true
		ret := dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
		if ret < 4 {
			grid[i][j] = color
		}
		return 1
	}
	dfs(r0, c0)
	return grid
}

// @lc code=end

