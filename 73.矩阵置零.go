/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])
	if m == 0 {
		return
	}
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, i*m+j)
			}
		}
	}
	for _, num := range queue {
		row := num / m
		col := num % m
		for j := 0; j < m; j++ {
			matrix[row][j] = 0
		}
		for i := 0; i < n; i++ {
			matrix[i][col] = 0
		}
	}
}

// @lc code=end

