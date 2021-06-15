/*
 * @lc app=leetcode.cn id=892 lang=golang
 *
 * [892] 三维形体的表面积
 */

// @lc code=start
func surfaceArea(grid [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	col := func(num int) int {
		// 0 2 4 6
		// return num*6 - (num-1)*2
		if num == 0 {
			return 0
		}
		return 4*num + 2
	}
	n := len(grid)
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans += col(grid[i][j])
		}
	}
	// row
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			ans -= min(grid[i][j], grid[i][j+1]) * 2
		}
	}
	// col
	for j := 0; j < n; j++ {
		for i := 0; i < n-1; i++ {
			ans -= min(grid[i][j], grid[i+1][j]) * 2
		}
	}
	return ans
}

// @lc code=end

