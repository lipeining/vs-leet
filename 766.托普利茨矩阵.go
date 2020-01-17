/*
 * @lc app=leetcode.cn id=766 lang=golang
 *
 * [766] 托普利茨矩阵
 */

// @lc code=start
func isToeplitzMatrix(matrix [][]int) bool {
	m:=len(matrix)
	n:=len(matrix[0])
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			if i>0&&j>0 {
				if matrix[i-1][j-1] != matrix[i][j] {
					return false
				}
			}
		}
	}
	return true
}
// @lc code=end

