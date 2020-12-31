/*
 * @lc app=leetcode.cn id=664 lang=golang
 *
 * [664] 奇怪的打印机
 */

// @lc code=start
func strangePrinter(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		dp[i][i] = 1
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1
			dp[i][j] = dp[i+1][j] + 1
			for k := i + 1; k <= j; k++ {
				if s[i] == s[k] {
					dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k+1][j])
				}
			}
		}
	}
	return dp[0][n-1]
	// n := len(s)
	// memo := make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	memo[i] = make([]int, n)
	// }
	// min := func(a, b int) int {
	// 	if a < b {
	// 		return a
	// 	}
	// 	return b
	// }
	// // 官方解答
	// var dp func(i, j int) int
	// dp = func(i, j int) int {
	// 	if i > j {
	// 		return 0
	// 	}
	// 	if memo[i][j] == 0 {
	// 		ans := dp(i+1, j) + 1
	// 		for k := i + 1; k <= j; k++ {
	// 			if s[k] == s[i] {
	// 				ans = min(ans, dp(i, k-1)+dp(k+1, j))
	// 			}
	// 		}
	// 		memo[i][j] = ans
	// 	}
	// 	return memo[i][j]
	// }
	// return dp(0, n-1)
}

// @lc code=end

