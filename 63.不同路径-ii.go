/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				if j >= 1 && obstacleGrid[i][j-1] == 0 {
					dp[j] += dp[j-1]
				}
			}
		}
	}
	return dp[n-1]
	// dp := make([][]int, m)
	// for i := 0; i < m; i++ {
	// 	dp[i] = make([]int, n)
	// }
	// // dp[0][0] = obstacleGrid[i][j]
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if i == 0 {
	// 			if obstacleGrid[i][j] == 1 {
	// 				dp[i][j] = 0
	// 			} else {
	// 				if j >= 1 {
	// 					dp[i][j] = dp[i][j-1]
	// 				} else {
	// 					dp[i][j] = 1
	// 				}
	// 			}
	// 		} else if j == 0 {
	// 			if obstacleGrid[i][j] == 1 {
	// 				dp[i][j] = 0
	// 			} else {
	// 				if i >= 1 {
	// 					dp[i][j] = dp[i-1][j]
	// 				} else {
	// 					dp[i][j] = 1
	// 				}
	// 			}
	// 		} else {
	// 			if obstacleGrid[i][j] == 1 {
	// 				dp[i][j] = 0
	// 			} else {
	// 				dp[i][j] = dp[i-1][j] + dp[i][j-1]
	// 			}
	// 		}
	// 	}
	// }
	// return dp[m-1][n-1]
}

// @lc code=end
