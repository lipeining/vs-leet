/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	if n == 0 {
		return ans
	}
	ans[0][0] = 1
	cnt := 2
	row := 0
	col := 1
	dir := "right"
	out := n-1
	inner := 0
	for cnt <= n*n {
		ans[row][col] = cnt
		cnt++
		if dir == "right" {
			if col == out {
				row++
				dir = "down"
			} else {
				col++
			}
		} else if dir == "down" {
			if row == out {
				col--
				dir = "left"
				out--
			} else {
				row++
			}
		} else if dir == "left" {
			if col == inner {
				row--
				dir = "up"
			} else {
				col--
			}
		} else if dir == "up" {
			if row == inner + 1 {
				col++
				dir = "right"
				inner++
			} else {
				row--
			}
		} else {
			// no way
		}

	}
	return ans
}

// @lc code=end

