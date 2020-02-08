import "math"

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
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

	// 其实是限制 K = infinity 的状况，可以当天卖出当天买入
	// k 和 k-1 是一样的
	// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k][0] - prices[i])
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
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
		tmp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, tmp-prices[i])
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
