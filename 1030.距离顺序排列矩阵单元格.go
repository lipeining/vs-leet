/*
 * @lc app=leetcode.cn id=1030 lang=golang
 *
 * [1030] 距离顺序排列矩阵单元格
 */

// @lc code=start
func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	ans := make([][]int, 0)
	queue := make([]int, 0)
	getIndex := func(x, y int) int {
		return x*cols + y
	}
	getXY := func(index int) (x, y int) {
		x = index / cols
		y = index % cols
		return
	}
	center := getIndex(rCenter, cCenter)
	queue = append(queue, center)
	seen := make([]bool, rows*cols)
	seen[center] = true
	// dis := 0
	dirs := [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			index := queue[i]
			x, y := getXY(index)
			ans = append(ans, []int{x, y})
			for _, dir := range dirs {
				nx, ny := x+dir[0], y+dir[1]
				if nx < 0 || nx == rows || ny < 0 || ny == cols {
					continue
				}
				nindex := getIndex(nx, ny)
				if seen[nindex] {
					continue
				}
				seen[nindex] = true
				queue = append(queue, nindex)
			}
		}
		queue = queue[size:]
	}
	return ans
}

// @lc code=end

