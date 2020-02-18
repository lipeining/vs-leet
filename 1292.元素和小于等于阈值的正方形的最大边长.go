/*
 * @lc app=leetcode.cn id=1292 lang=golang
 *
 * [1292] 元素和小于等于阈值的正方形的最大边长
 */

// @lc code=start
func maxSideLength(mat [][]int, threshold int) int {
	m := len(mat)
	n := len(mat[0])
	p := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		p[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			p[i][j] = p[i-1][j] + p[i][j-1] - p[i-1][j-1] + mat[i-1][j-1]
		}
	}
	left := 1
	right := min(m, n)
	ans := 0
	for left <= right {
		mid := left + (right-left)>>1
		find := false
		for i := 1; i < m-mid+1; i++ {
			for j := 1; j <= n-mid+1; j++ {
				if getSum(p, i, j, i+mid-1, j+mid-1) <= threshold {
					find = true
				}
			}
		}
		if find {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
func getSum(p [][]int, x1, y1, x2, y2 int) int {
	return p[x2][y2] - p[x1-1][y2] - p[x2][y1-1] + p[x1-1][y1-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
