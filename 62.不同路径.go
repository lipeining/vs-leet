/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	// 压缩路径
	dp := make([]int, n)
	for j := 0; j < n; j++ {
		dp[j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]

	// dp := make([][]int, m)
	// for i := 0; i < m; i++ {
	// 	dp[i] = make([]int, n)
	// }
	// dp[0][0] = 1
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if i == 0 {
	// 			dp[i][j] = 1
	// 		} else if j == 0 {
	// 			dp[i][j] = 1
	// 		} else {
	// 			dp[i][j] = dp[i-1][j] + dp[i][j-1]
	// 		}
	// 	}
	// }
	// return dp[m-1][n-1]
}

// @lc code=end
