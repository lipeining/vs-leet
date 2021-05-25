/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	if cols == 0 {
		return 0
	}
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
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
	// 以 i,j 为底的边长
	ans := 0
	// 区间 dp 枚举高度
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			dp[i][j] = 1
			if j >= 1 {
				dp[i][j] = dp[i][j-1] + 1
			}
			width := dp[i][j]
			for k := i; k >= 0; k-- {
				width = min(width, dp[k][j])
				ans = max(ans, width*(i-k+1))
			}
		}
	}
	return ans
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end

