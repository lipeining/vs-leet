/*
 * @lc app=leetcode.cn id=787 lang=golang
 *
 * [787] K 站中转内最便宜的航班
 */

// @lc code=start
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp[src][0] = 0
	for _, flight := range flights {
		start := flight[0]
		end := flight[1]
		price := flight[2]
		if start == src {
			dp[end][0] = min(dp[end][0], price)
		}
	}
	for j := 1; j <= K; j++ {
		for _, flight := range flights {
			start := flight[0]
			end := flight[1]
			price := flight[2]
			dp[end][j] = min(dp[end][j-1], min(dp[end][j], dp[start][j-1]+price))
		}
	}
	if dp[dst][K] == math.MaxInt32 {
		return -1
	}
	return dp[dst][K]
}

// @lc code=end

