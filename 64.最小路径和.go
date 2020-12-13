/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j >= 1 {
					dp[j] = dp[j-1] + grid[i][j]
				}
			} else if j == 0 {
				if i >= 1 {
					dp[j] = dp[j] + grid[i][j]
				}
			} else {
				dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
	// dp := make([][]int, m)
	// for i := 0; i < m; i++ {
	// 	dp[i] = make([]int, n)
	// }
	// dp[0][0] = grid[0][0]
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if i == 0 {
	// 			if j >= 1 {
	// 				dp[i][j] = dp[i][j-1] + grid[i][j]
	// 			}
	// 		} else if j == 0 {
	// 			if i >= 1 {
	// 				dp[i][j] = dp[i-1][j] + grid[i][j]
	// 			}
	// 		} else {
	// 			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
	// 		}
	// 	}
	// }
	// return dp[m-1][n-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
