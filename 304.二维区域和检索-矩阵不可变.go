/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start
type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 {
				if j >= 1 {
					matrix[i][j] = matrix[i][j-1] + matrix[i][j]
				}
			} else if j == 0 {
				if i >= 1 {
					matrix[i][j] = matrix[i-1][j] + matrix[i][j]
				}
			} else {
				matrix[i][j] = matrix[i-1][j] + matrix[i][j-1] + matrix[i][j] - matrix[i-1][j-1]
			}
		}
	}
	return NumMatrix{sum: matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := this.sum
	ans := sum[row2][col2]
	if row1 == 0 && col1 == 0 {
		return ans
	} else if row1 == 0 {
		return ans - sum[row2][col1-1]
	} else if col1 == 0 {
		return ans - sum[row1-1][col2]
	} else {
		return ans + sum[row1-1][col1-1] - sum[row2][col1-1] - sum[row1-1][col2]
	}
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end
