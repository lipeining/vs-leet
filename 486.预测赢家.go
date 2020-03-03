/*
 * @lc app=leetcode.cn id=486 lang=golang
 *
 * [486] 预测赢家
 */

// @lc code=start
func PredictTheWinner(nums []int) bool {
	// /dp(i, j) = max(nums[i] - dp(i+1, j), nums[j] - dp(i, j-1))
	// dp(i, i) = nums[i]
	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n)
	}
	for i := n; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			a := nums[i] - dp[i+1][j]
			b := nums[j] - dp[i][j-1]
			dp[i][j] = max(a, b)
		}
	}
	return dp[0][n-1] >= 0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
