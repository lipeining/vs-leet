/*
 * @lc app=leetcode.cn id=1074 lang=golang
 *
 * [1074] 元素和为目标值的子矩阵数量
 */

// @lc code=start
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	n := len(matrix)
	m := len(matrix[0])
	pre := make([][]int, n)
	for i := 0; i < n; i++ {
		pre[i] = make([]int, m)
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				pre[i][j] = matrix[i][j]
			} else if i == 0 {
				pre[i][j] = pre[i][j-1] + matrix[i][j]
			} else if j == 0 {
				pre[i][j] = pre[i-1][j] + matrix[i][j]
			} else {
				pre[i][j] = pre[i][j-1] + pre[i-1][j] + matrix[i][j] - pre[i-1][j-1]
			}
			for x := 0; x <= i; x++ {
				for y := 0; y <= j; y++ {
					c := 0
					if x > 0 && y > 0 {
						c = pre[x-1][y-1]
					}
					a := 0
					if x > 0 {
						a = pre[x-1][y]
					}
					b := 0
					if y > 0 {
						b = pre[x][y-1]
					}
					v := pre[i][j] - a - b + c
					if v == target {
						count++
					}
				}
			}
		}
	}
	return count
}

// @lc code=end

