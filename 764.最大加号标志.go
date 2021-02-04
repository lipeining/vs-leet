/*
 * @lc app=leetcode.cn id=764 lang=golang
 *
 * [764] 最大加号标志
 */

// @lc code=start
func orderOfLargestPlusSign(N int, mines [][]int) int {
	banned := make(map[int]bool)
	for i := 0; i < len(mines); i++ {
		index := N*mines[i][0] + mines[i][1]
		banned[index] = true
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
	}
	ans := 0
	count := 0
	for i := 0; i < N; i++ {
		count = 0
		for j := 0; j < N; j++ {
			index := i*N + j
			if ban := banned[index]; !ban {
				count = count + 1
			} else {
				count = 0
			}
			dp[i][j] = count
		}
		count = 0
		for j := N-1; j >= 0; j-- {
			index := i*N + j
			if ban := banned[index]; !ban {
				count = count + 1
			} else {
				count = 0
			}
			dp[i][j] = min(dp[i][j], count)
		}
	}
	for j := 0; j < N; j++ {
		count = 0
		for i := 0; i < N; i++ {
			index := i*N + j
			if ban := banned[index]; !ban {
				count = count + 1
			} else {
				count = 0
			}
			dp[i][j] = min(dp[i][j], count)
		}
		count = 0
		for i := N-1; i >= 0; i-- {
			index := i*N + j
			if ban := banned[index]; !ban {
				count = count + 1
			} else {
				count = 0
			}
			dp[i][j] = min(dp[i][j], count)
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}

// @lc code=end

