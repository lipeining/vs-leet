/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]

	// if n <= 1 {
	// 	return n
	// }
	// dp := make([]int, n+1)
	// // dp[0] = 0
	// for j := 1; j <= n; j++ {
	// 	if j*j <= n {
	// 		dp[j*j] = 1
	// 	}
	// 	if dp[j] == 0 {
	// 		dp[j] = n
	// 	}
	// }
	// // dp[i] = minN(dp[i-k]+dp[k])
	// for i := 1; i <= n; i++ {
	// 	should := n
	// 	for j := 0; j <= i/2; j++ {
	// 		should = min(should, dp[i-j]+dp[j])
	// 	}
	// 	dp[i] = should
	// 	// fmt.Println(i, j, should)
	// }
	// return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
