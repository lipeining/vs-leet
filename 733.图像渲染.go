/*
 * @lc app=leetcode.cn id=733 lang=golang
 *
 * [733] 图像渲染
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	m := len(image)
	n := len(image[0])
	seen := make([][]bool, m)
	for i := 0; i < m; i++ {
		seen[i] = make([]bool, n)
	}
	to := image[sr][sc]
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	var dfs func(i, j int, to int)
	dfs = func(i, j int, to int) {
		if i < 0 || i == m || j < 0 || j == n {
			return
		}
		if seen[i][j] {
			return
		}
		if image[i][j] != to {
			return
		}
		seen[i][j] = true
		image[i][j] = newColor
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			dfs(ni, nj, to)
		}
	}
	dfs(sr, sc, to)
	return image
}

// @lc code=end

