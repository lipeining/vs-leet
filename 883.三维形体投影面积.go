/*
 * @lc app=leetcode.cn id=883 lang=golang
 *
 * [883] 三维形体投影面积
 */

// @lc code=start
func projectionArea(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return n
	}
	// 按行，按列取最大值
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				ans++
			}
		}
	}
	for i := 0; i < n; i++ {
		col := grid[i][0]
		for j := 0; j < n; j++ {
			col = max(col, grid[i][j])
		}
		ans += col
	}
	for j := 0; j < n; j++ {
		row := grid[0][j]
		for i := 0; i < n; i++ {
			row = max(row, grid[i][j])
		}
		ans += row
	}
	return ans
}

// @lc code=end

