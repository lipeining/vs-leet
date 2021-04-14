// func gcd(x, y int)int{

// }
func maxScore(nums []int) int {
	var gcd func(x, y int) int
	gcd = func(x, y int) int {
		if y == 0 {
			return x
		}
		return gcd(y, x%y)
	}
	n := len(nums)
	dp := make([]int, 1<<n)
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			g[i][j] = gcd(nums[i], nums[j])
		}
	}
	popcount := func(x int) int {
		cnt := 0
		for x != 0 {
			x &= x - 1
			cnt++
		}
		return cnt
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for state := 0; state < (1 << n); state++ {
		cnt := popcount(state)
		if cnt&1 == 1 {
			// 只有一个 1 ，需要成对出现
			continue
		}
		// now=计算当前是第几次 gcd
		now := cnt/2 + 1
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if (state&(1<<i) == 1) || (state&(1<<j) == 1) {
					// i, j 的位置已经被占用了
					continue
				}
				nxt := state | (1 << i) | (1 << j)
				dp[nxt] = max(dp[nxt], dp[state]+g[i][j]*now)
			}
		}
	}
	return dp[1<<n-1]
}