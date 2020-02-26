/*
 * @lc app=leetcode.cn id=650 lang=golang
 *
 * [650] 只有两个键的键盘
 */

// @lc code=start
func minSteps(n int) int {
	if n <= 1 {
		return 0
	}
	dp := make([]int, n+1)
	// 全是 paste
	for i := 1; i <= n; i++ {
		dp[i] = i
	}
	// 优化为 j:=i/2 j>=2 即可， 可以提前跳出 i/%j == 0
	for i := 1; i <= n; i++ {
		// 从两种操作中选一个
		// 比如： i=6  时，
		// 可以通过  dp[2]+3 || dp[3]+2 的方法得到 dp[6]
		for j := 2; j < i; j++ {
			if i%j == 0 {
				dp[i] = min(dp[i], dp[j]+i/j)
			} else {
				// do nothing
			}
		}
	}
	// fmt.Println(dp)
	return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
