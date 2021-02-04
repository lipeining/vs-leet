/*
 * @lc app=leetcode.cn id=1139 lang=golang
 *
 * [1139] 最大的以 1 为边界的正方形
 */

// @lc code=start
func largest1BorderedSquare(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if grid[i-1][j-1] == 1 {
				dp[i][j][0] = 1 + dp[i][j-1][0]
				dp[i][j][1] = 1 + dp[i-1][j][1]
			}
		}
	}
	ans := 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for side := min(dp[i][j][0], dp[i][j][1]); side >= 1; side-- {
				if dp[i][j-side+1][1] >= side && dp[i-side+1][j][0] >= side {
					ans = max(ans, side)
					break
				}
			}
		}
	}
	return ans * ans
}

// @lc code=end

