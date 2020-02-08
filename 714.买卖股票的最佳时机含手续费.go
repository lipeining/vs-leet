import "math"

/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	// base case：
	// dp[-1][k][0] = dp[i][0][0] = 0
	// dp[-1][k][1] = dp[i][0][1] = -infinity

	// 状态转移方程：
	// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

	// 需要手续费，所以，卖的时候需要减去手续费
	// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i] - fee)

	// k = infinity
	// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k][0] - prices[i] - fee)
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i] - fee)
	// base case :
	// dp[-1][0] = dp[i][0] = 0
	// dp[-1][1] = dp[i][1] = -infinity

	length := len(prices)
	if length <= 1 {
		return 0
	}
	dp_i_0 := 0
	dp_i_1 := math.MinInt32
	for i := 0; i < length; i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, dp_i_0-prices[i]-fee)
	}
	return dp_i_0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
