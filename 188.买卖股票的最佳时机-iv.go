import "math"

/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	maxK := k
	length := len(prices)
	if maxK > length/2 {
		return maxProfitInf(prices)
	}

	// base case：
	// dp[-1][k][0] = dp[i][0][0] = 0
	// dp[-1][k][1] = dp[i][0][1] = -infinity

	// 状态转移方程：
	// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
	dp := make([][][]int, length)
	state := 2
	for i := 0; i < length; i++ {
		dp[i] = make([][]int, maxK+1)
		for j := 0; j <= maxK; j++ {
			dp[i][j] = make([]int, state)
		}
	}
	for i := 0; i < length; i++ {
		for j := 1; j <= maxK ; j++ {
			if i == 0 {
				dp[i][j][0] = 0
				// dp[i][j][1] = math.MinInt32
				// 这里应该是 -prices[i]
				dp[i][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
		// for j := maxK; j >= 1; j-- {
		// 	if i == 0 {
		// 		dp[i][j][0] = 0
		// 		// dp[i][j][1] = math.MinInt32
		// 		// 这里应该是 -prices[i]
		// 		dp[i][j][1] = -prices[i]
		// 		continue
		// 	}
		// 	dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
		// 	dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		// }
	}
	return dp[length-1][maxK][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfitInf(prices []int) int {
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

// @lc code=end
