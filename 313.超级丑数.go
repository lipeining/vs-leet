/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */

// @lc code=start
func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n)
	dp[0] = 1
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	np := len(primes)
	p := make([]int, np)
	for i := 1; i < n; i++ {
		curMin := math.MaxInt32
		for j := 0; j < np; j++ {
			curMin = min(curMin, primes[j]*dp[p[j]])
		}
		dp[i] = curMin
		for j := 0; j < np; j++ {
			if curMin == primes[j]*dp[p[j]] {
				p[j]++
			}
		}
	}
	return dp[n-1]
}

// @lc code=end

