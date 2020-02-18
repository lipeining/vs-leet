/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	left, right := 0, m*n-1
	for left < right {
		mid := left + (right-left)>>1
		now := helper(matrix, m, n, mid)
		// fmt.Println(mid, now)
		if now == target {
			return true
		} else if now > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// fmt.Println(left, right)
	if helper(matrix, m, n, left) == target {
		return true
	}
	return false
}
func helper(matrix [][]int, m, n, current int) int {
	i := current / n
	j := current % n
	return matrix[i][j]
}

// @lc code=end
