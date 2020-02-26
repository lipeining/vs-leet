/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	// 右下角判断方法
	// 以右下角为边界，画三个小一点的正方形
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	// dp[i][j] = dp[i-1][j] dp[i][j-1] dp[i-1][j-1]
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = int(matrix[i][j] - '0')
			} else {
				if matrix[i][j] == '0' {
					dp[i][j] = 0
				} else {
					dp[i][j] = min3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				}
			}
			if dp[i][j]*dp[i][j] > ans {
				ans = dp[i][j] * dp[i][j]
			}
		}
	}
	// fmt.Println(dp)
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func min3(a, b, c int) int {
	return min(a, min(b, c))
}

// @lc code=end
