/*
 * @lc app=leetcode.cn id=1223 lang=golang
 *
 * [1223] 掷骰子模拟
 */

// @lc code=start
func dieSimulator(n int, rollMax []int) int {
	// dp[step][i][j]
	// 第 step 次投掷时，当前为 i，连续 j 次
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 7)
		for j := 0; j < 7; j++ {
			dp[i][j] = make([]int, 16)
		}
	}
	mod := int(1e9 + 7)
	for step := 1; step <= n; step++ {
		for i := 1; i <= 6; i++ {
			if step == 1 {
				dp[step][i][1] = 1
			} else {
				for j := 2; j <= rollMax[i-1]; j++ {
					dp[step][i][j] = dp[step-1][i][j-1]
				}
				// 其他不是 i 的
				sum := 0
				for t := 1; t <= 6; t++ {
					if t != i {
						for tj := 1; tj <= 15; tj++ {
							sum = (sum + dp[step-1][t][tj]) % mod
						}
					}
				}
				dp[step][i][1] = sum
			}
		}
	}
	ans := 0
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 15; j++ {
			ans = (ans + dp[n][i][j]) % mod
		}
	}
	return ans
}

// @lc code=end
