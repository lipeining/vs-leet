

func minTaps(n int, ranges []int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = maht.MaxInt32
	}
	dp[0] = 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i <= n; i++ {
		left := max(0, i-ranges[i])
		right := min(n, i+ranges[i])
		for j := left + 1; j <= right; j++ {
			if dp[left] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[left]+1)
			}
		}
	}
	if dp[n] == math.MaxInt32 {
		return -1
	}
	return dp[n]
}