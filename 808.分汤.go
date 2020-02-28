/*
 * @lc app=leetcode.cn id=808 lang=golang
 *
 * [808] 分汤
 */

// @lc code=start
func soupServings(N int) float64 {
	mod := N % 25
	N = N / 25
	if mod != 0 {
		N++
	}
	dp := make([][]float64, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]float64, N+1)
	}
	// 官方答案
	for s := 2; s <= 2*N; s++ {
		for i := 0; i <= N; i++ {
			j := s - i
			if j < 0 || j > N {
				continue
			}
			ans := 0.0
			if i == 0 {
				ans = 1.0
			}
			if i == 0 && j == 0 {
				ans = 0.5
			}
			if i > 0 && j > 0 {
				ans = 0.25 * (dp[M(i-4)][j] + dp[M(i-3)][M(j-1)] + dp[M(i-2)][M(j-2)] + dp[M(i-1)][M(j-3)])
			}
			dp[i][j] = ans
		}
	}
	return dp[N][N]
}
func M(b int) int {
	return max(0, b)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
