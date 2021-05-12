/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}
	seen := make([][]bool, m)
	for i := 0; i < m; i++ {
		seen[i] = make([]bool, n)
	}
	type pair struct {
		i int
		j int
	}
	queue := make([]pair, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, pair{i, j})
				seen[i][j] = true
			}
		}
	}
	dirs := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for len(queue) != 0 {
		size := len(queue)
		for t := 0; t < size; t++ {
			p := queue[t]
			for _, dir := range dirs {
				ni := p.i + dir[0]
				nj := p.j + dir[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n && !seen[ni][nj] {
					ans[ni][nj] = ans[p.i][p.j] + 1
					queue = append(queue, pair{i: ni, j: nj})
					seen[ni][nj] = true
				}
			}
		}
		queue = queue[size:]
	}
	return ans
}
func updateMatrixw(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != 0 {
				ans[i][j] = math.MaxInt32
			}
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	valid := func(x, y int) bool {
		return !(x < 0 || x == m || y < 0 || y == n)
	}
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if ans[x][y] != math.MaxInt32 {
			return ans[x][y]
		}
		left, right, up, down := math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32
		if valid(x, y-1) {
			left = dfs(x, y-1)
		}
		if valid(x, y+1) {
			right = dfs(x, y+1)
		}
		if valid(x-1, y) {
			up = dfs(x-1, y)
		}
		if valid(x+1, y) {
			down = dfs(x+1, y)
		}
		ret := min(min(min(left, right), up), down)
		ans[x][y] = ret + 1
		return ans[x][y]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ans[i][j] != 0 {
				dfs(i, j)
			}
		}
	}
	return ans
}

// @lc code=end

