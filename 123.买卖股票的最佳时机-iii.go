import "math"

/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	// base case：
	// dp[-1][k][0] = dp[i][0][0] = 0
	// dp[-1][k][1] = dp[i][0][1] = -infinity

	// 状态转移方程：
	// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
	// 状态转移方程：

	// 其实是限制 K = 2 的状况，可以当天卖出当天买入
	// dp_i_1_0 = max(dp_i_1_0, dp_i_1_1+prices[i])
	// dp_i_2_0 = max(dp_i_2_0, dp_i_2_1+prices[i])
	// dp_i_1_1 = max(dp_i_1_1, dp_i_0_0-prices[i])
	// dp_i_2_1 = max(dp_i_2_1, dp_i_1_0-prices[i])

	// dp_i_0_0 = 0
	// dp_i_0_1 = -infinity

	length := len(prices)
	if length <= 1 {
		return 0
	}
	dp_i_0_0 := 0
	dp_i_1_0 := 0
	dp_i_1_1 := math.MinInt32
	dp_i_2_0 := 0
	dp_i_2_1 := math.MinInt32
	for i := 0; i < length; i++ {
		// other
		dp_i_2_0 = max(dp_i_2_0, dp_i_2_1+prices[i])
		dp_i_2_1 = max(dp_i_2_1, dp_i_1_0-prices[i])
		dp_i_1_0 = max(dp_i_1_0, dp_i_1_1+prices[i])
		dp_i_1_1 = max(dp_i_1_1, dp_i_0_0-prices[i])

		// own
		// dp_i_1_0 = max(dp_i_1_0, dp_i_1_1+prices[i])
		// dp_i_1_1 = max(dp_i_1_1, dp_i_0_0-prices[i])
		// dp_i_2_0 = max(dp_i_2_0, dp_i_2_1+prices[i])
		// dp_i_2_1 = max(dp_i_2_1, dp_i_1_0-prices[i])
	}
	// 交换 K 的遍历顺序是不影响的
	return dp_i_2_0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
