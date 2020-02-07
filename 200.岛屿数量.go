/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	var dfs func(grid [][]byte, x int, y int)
	dfs = func(grid [][]byte, x int, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		dfs(grid, x-1, y)
		dfs(grid, x+1, y)
		dfs(grid, x, y-1)
		dfs(grid, x, y+1)
	}
	counter := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				counter++
				dfs(grid, i, j)
			}
		}
	}
	return counter
}

// @lc code=end
