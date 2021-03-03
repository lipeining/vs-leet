/*
 * @lc app=leetcode.cn id=1277 lang=golang
 *
 * [1277] 统计全为 1 的正方形子矩阵
 */

// @lc code=start
func countSquares(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}
	ans := 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				f[i][j] = matrix[i][j]
			} else if matrix[i][j] == 0 {
				f[i][j] = 0
			} else {
				f[i][j] = min(min(f[i-1][j-1], f[i-1][j]), f[i][j-1]) + 1
			}
			ans += f[i][j]
		}
	}
	return ans
}

// @lc code=end

