/*
 * @lc app=leetcode.cn id=920 lang=golang
 *
 * [920] 播放列表的数量
 */

// @lc code=start
func numMusicPlaylists(N int, L int, K int) int {
	dp := make([][]int, L+1)
	for i := 0; i <= L; i++ {
		dp[i] = make([]int, N+1)
	}
	dp[0][0] = 1
	toMod := int(1e9 + 7)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i <= L; i++ {
		for j := 1; j <= N; j++ {
			dp[i][j] += dp[i-1][j-1] * (N - j + 1)
			dp[i][j] += dp[i-1][j] * max(j-K, 0)
			dp[i][j] %= toMod
		}
	}
	return dp[L][N]
}

// @lc code=end

