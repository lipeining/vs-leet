/*
 * @lc app=leetcode.cn id=1220 lang=golang
 *
 * [1220] 统计元音字母序列的数目
 */

// @lc code=start
func countVowelPermutation(n int) int {
	dp := make([][]int, 5)
	for i := 0; i < 5; i++ {
		dp[i] = make([]int, n+1)
	}
	// dp[0][n-1]
	mod := int(1e9 + 7)
	for step := 1; step <= n; step++ {
		for i := 0; i < 5; i++ {
			if step == 1 {
				dp[i][step] = 1
			} else {
				// 0 1 2 3 4
				// a e i o u
				if i == 0 {
					// 前面只能是 e, i, u
					dp[i][step] = (dp[1][step-1] + dp[2][step-1] + dp[4][step-1]) % mod
				} else if i == 1 {
					// a i
					dp[i][step] = (dp[0][step-1] + dp[2][step-1]) % mod
				} else if i == 2 {
					// e, o,
					dp[i][step] = (dp[1][step-1] + dp[3][step-1]) % mod
				} else if i == 3 {
					// i
					dp[i][step] = dp[2][step-1] % mod
				} else {
					// i o
					dp[i][step] = (dp[2][step-1] + dp[3][step-1]) % mod
				}
			}
		}
	}
	ans := 0
	for i := 0; i < 5; i++ {
		ans = (ans + dp[i][n]) % mod
	}
	return ans
}

// @lc code=end
