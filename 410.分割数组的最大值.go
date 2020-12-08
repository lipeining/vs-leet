/*
 * @lc app=leetcode.cn id=410 lang=golang
 *
 * [410] 分割数组的最大值
 */

// @lc code=start
func splitArray(nums []int, m int) int {
	n := len(nums)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + nums[i]
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	// 前 i 个数分为 j 个区间
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, m); j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = min(dp[i][j], max(dp[k][j-1], pre[i]-pre[k]))
			}
		}
	}
	return dp[n][m]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

