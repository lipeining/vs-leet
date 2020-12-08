/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)
	// dp[0] = 1
	// dp[1] = 1
	// dp[2] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			// dp[i] = max(dp[i], max3(dp[i-j]*dp[j], dp[i-j]*j, (i-j)*j))
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
		// should := dp[i]
		// for j := 1; j <= i; j++ {
		// 	should = max(should, dp[i-j]*dp[j])
		// }
		// dp[i] = should
	}
	// fmt.Println(dp)
	return dp[n]
}
func max3(a, b, c int) int {
	return max(a, max(b, c))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
