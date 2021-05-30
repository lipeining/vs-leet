/*
 * @lc app=leetcode.cn id=1001 lang=golang
 *
 * [1001] 网格照明
 */
package main

// func main() {
// 	gridIllumination(5, [][]int{{0, 0}, {4, 4}}, [][]int{{1, 1}, {1, 0}})
// 	gridIllumination(5, [][]int{{0, 0}, {4, 4}}, [][]int{{1, 1}, {1, 1}})
// 	gridIllumination(5, [][]int{{0, 0}, {0, 4}}, [][]int{{0, 4}, {0, 1}, {1, 4}})
// }

// @lc code=start
func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	parseIndex := func(num int) (int, int) {
		return num / n, num % n
	}
	toIndex := func(x, y int) int {
		return x*n + y
	}
	lm := make(map[int]bool)
	for _, lamp := range lamps {
		lm[toIndex(lamp[0], lamp[1])] = true
	}
	ans := make([]int, len(queries))
	dirs := [][]int{
		{0, 0},
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	// 应该改用四个方向的 hash map
	// 判断是否符合即可，
	// 举例，第 i 行有三盏灯，删除时，减一。
	light := func(row, col int) int {
		for k := range lm {
			i, j := parseIndex(k)
			if i == row {
				return 1
			}
			if j == col {
				return 1
			}
			if abs(i-row) == abs(j-col) {
				return 1
			}
		}
		return 0
	}
	down := func(row, col int) {
		// fmt.Println("before", lm)
		for _, dir := range dirs {
			x, y := row+dir[0], col+dir[1]
			if x < 0 || x >= n || y < 0 || y >= n {
				continue
			}
			index := toIndex(x, y)
			// fmt.Println("on", row, col, "delete", x, y)
			delete(lm, index)
		}
		// fmt.Println("after", lm)
	}
	for i := 0; i < len(queries); i++ {
		row, col := queries[i][0], queries[i][1]
		ans[i] = light(row, col)
		down(row, col)
	}
	// fmt.Println("ansssssssssss", ans)
	// fmt.Println(lm)
	return ans
}

// @lc code=end
