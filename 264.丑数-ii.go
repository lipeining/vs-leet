/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	// 抄思路
	dp := make([]int, n+1)
	dp[1] = 1
	index2, index3, index5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = min3(dp[index2]*2, dp[index3]*3, dp[index5]*5)
		for dp[index2]*2 <= dp[i] {
			index2++
		}
		for dp[index3]*3 <= dp[i] {
			index3++
		}
		for dp[index5]*5 <= dp[i] {
			index5++
		}
	}
	return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func min3(a, b, c int) int {
	return min(a, min(b, c))
}

// @lc code=end
