import (
	"math"
)

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-l-3/
	// base case：
	// dp[-1][k][0] = dp[i][0][0] = 0
	// dp[-1][k][1] = dp[i][0][1] = -infinity

	// 状态转移方程：
	// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

	// dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
	// dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0] - prices[i])
	//             = max(dp[i-1][1][1], -prices[i])
	// 解释：k = 0 的 base case，所以 dp[i-1][0][0] = 0。

	// 现在发现 k 都是 1，不会改变，即 k 对状态转移已经没有影响了。
	// 可以进行进一步化简去掉所有 k：
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
	// dp[i][1] = max(dp[i-1][1], -prices[i])

	// length := len(prices)
	// if length <= 1 {
	// 	return 0
	// }
	// K := 1
	// state := 2
	// dp := make([][][]int, length)
	// for i := 0; i < length; i++ {
	// 	dp[i] = make([][]int, K+1)
	// 	for j := 0; j <= K; j++ {
	// 		dp[i][j] = make([]int, state)
	// 	}
	// 	dp[i][0][0] = 0
	// 	dp[i][0][1] = math.MinInt32
	// }
	// // fmt.Println(dp)
	// for i := 0; i < length; i++ {
	// 	for k := 1; k <= K; k++ {
	// 		if i == 0 {
	// 			dp[i][k][0] = 0
	// 			dp[i][k][1] = -prices[i]
	// 			continue
	// 		}
	// 		dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
	// 		dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
	// 	}
	// }
	// // fmt.Println(dp)
	// return dp[length-1][K][0]

	// length := len(prices)
	// if length <= 1 {
	// 	return 0
	// }
	// state := 2
	// dp := make([][]int, length)
	// for i := 0; i < length; i++ {
	// 	dp[i] = make([]int, state)
	// 	dp[i][0] = 0
	// 	dp[i][1] = math.MinInt32
	// }
	// // fmt.Println(dp)
	// for i := 0; i < length; i++ {
	// 	if i == 0 {
	// 		dp[i][0] = 0
	// 		dp[i][1] = -prices[i]
	// 		continue
	// 	}
	// 	dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	// 	dp[i][1] = max(dp[i-1][1], -prices[i])
	// }
	// // fmt.Println(dp)
	// return dp[length-1][0]

	length := len(prices)
	if length <= 1 {
		return 0
	}
	dp_i_0 := 0
	dp_i_1 := math.MinInt32
	for i := 0; i < length; i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
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

// 画图理解的峰点，谷底
// 在 dp[i][j] 时
// length := len(prices)
// if length == 0 {
// 	return 0
// }
// min := math.MaxInt32
// max := 0
// for i := 0; i < length; i++ {
// 	if min > prices[i] {
// 		min = prices[i]
// 	}
// 	if max < prices[i]-min {
// 		max = prices[i] - min
// 	}
// }
// return max
// dp := make([][]int, length)
// for i := 0; i < length; i++ {
// 	dp[i] = make([]int, length)
// }
// dp[i] = max(0,dp[i-1])

// 简答的 dp
// length := len(prices)
// if length <= 1 {
// 	return 0
// }
// diff := make([]int, length)
// for i := 0; i < length-1; i++ {
// 	diff[i] = prices[i+1] - prices[i]
// }
// dp := make([]int, length)
// dp[0] = diff[0]
// if dp[0] < 0 {
// 	dp[0] = 0
// }
// profit := diff[0]
// for i := 1; i < length; i++ {
// 	max := dp[i-1] + diff[i]
// 	if max < 0 {
// 		max = 0
// 	}
// 	dp[i] = max
// 	if profit < max {
// 		profit = max
// 	}
// }
// return profit