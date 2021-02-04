/*
 * @lc app=leetcode.cn id=1240 lang=golang
 *
 * [1240] 铺瓷砖
 */

// @lc code=start
func tilingRectangle(n int, m int) int {
	// 他人答案
	
	// dp := make([][]int, n+1)
	// for i := 0; i <= n; i++ {
	// 	dp[i] = make([]int, m+1)
	// }
	// min := func(a, b int) int {
	// 	if a < b {
	// 		return a
	// 	}
	// 	return b
	// }
	// for i := 1; i <= n; i++ {
	// 	for j := 1; j <= m; j++ {
	// 		if i == j {
	// 			dp[i][j] = 1
	// 		} else {
	// 			side := min(i, j)
	// 			for k := side; k >= 1; k-- {

	// 			}
	// 		}
	// 	}
	// }
	// return dp[n][m]
}

// @lc code=end

