package fall

func keyboard(k int, n int) int {
	// 动态规划
	// dp[n][k] = dp[n][k-1] *
	// base case：
	// dp[-1][k][0] = dp[i][0][0] = 0
	// dp[-1][k][1] = dp[i][0][1] = -infinity

	// 状态转移方程： [][][1-26]
	// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
	if k >= n {
		return nTime(26, n)
	}
	if n == 0 {
		return 0
	}
	if k == 0 {
		return 0
	}
	// 如果 k < n 那么重复的需要进行删除操作。
	return 0
}
func nTime(num int, n int) int {
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= num
	}
	return ans
}

// func main() {

// }
