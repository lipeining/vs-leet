/*
 * @lc app=leetcode.cn id=940 lang=golang
 *
 * [940] 不同的子序列 II
 */

// @lc code=start
func distinctSubseqII(S string) int {
	last := make([]int, 26)
	for i := 0; i < 26; i++ {
		last[i] = -1
	}
	n := len(S)
	dp := make([]int, n+1)
	mod := int(1e9 + 7)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		x := S[i-1] - 'a'
		dp[i] = dp[i-1] * 2 % mod
		if last[x] >= 0 {
			dp[i] -= dp[last[x]]
		}
		dp[i] %= mod
		last[x] = i-1
	}
	dp[n]--
	if dp[n] < 0 {
		dp[n] += mod
	}
	return dp[n]
	// last := make([]int, 26)
	// for i := 0; i < 26; i++ {
	// 	last[i] = -1
	// }
	// n := len(S)
	// dp := make([]int, n+1)
	// mod := int(1e9 + 7)
	// dp[0] = 1
	// for i := 0; i < n; i++ {
	// 	x := S[i] - 'a'
	// 	dp[i+1] = dp[i] * 2 % mod
	// 	if last[x] >= 0 {
	// 		dp[i+1] -= dp[last[x]]
	// 	}
	// 	dp[i+1] %= mod
	// 	last[x] = i
	// }
	// dp[n]--
	// if dp[n] < 0 {
	// 	dp[n] += mod
	// }
	// return dp[n]
}

// @lc code=end

