/*
 * @lc app=leetcode.cn id=885 lang=golang
 *
 * [885] 螺旋矩阵 III
 */

// @lc code=start
func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	total := rows * cols
	ans := make([][]int, 0)
	dir := 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	inbox := func(i, j int) bool {
		return i >= 0 && i < rows && j >= 0 && j < cols
	}
	left := cStart - 1
	right := cStart + 1
	up := rStart - 1
	down := rStart + 1
	i, j := rStart, cStart
	// 1，1，2，2，3，3，4，4
	for len(ans) != total {
		if inbox(i, j) {
			ans = append(ans, []int{i, j})
		}
		if dir == 0 {
			if j == right {
				dir++
				right++
			}
		} else if dir == 1 {
			if i == down {
				dir++
				down++
			}
		} else if dir == 2 {
			if j == left {
				dir++
				left--
			}
		} else if dir == 3 {
			if i == up {
				dir=0
				up--
			}
		} else {
			// not happen
		}
		i += dirs[dir][0]
		j += dirs[dir][1]
	}
	return ans
}

// @lc code=end

