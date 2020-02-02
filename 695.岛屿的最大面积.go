/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	m := len(grid)
	if m == 0 {
		return ans
	}
	n := len(grid[0])
	if n == 0 {
		return ans
	}
	var dfs func(grid [][]int, x int, y int, seen [][]bool) int
	dfs = func(grid [][]int, x int, y int, seen [][]bool) int {
		if x < 0 || x >= m || y < 0 || y >= n {
			return 0
		}
		if grid[x][y] == 0 {
			return 0
		}
		if seen[x][y] {
			return 0
		}
		seen[x][y] = true
		return 1 + dfs(grid, x-1, y, seen) + dfs(grid, x+1, y, seen) + dfs(grid, x, y-1, seen) + dfs(grid, x, y+1, seen)	
	}
	seen := make([][]bool, m)
	for i:=0;i<m;i++ {
		seen[i] = make([]bool,n)
	}
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			if grid[i][j] == 1 {
				tmp := dfs(grid, i, j, seen)
				if tmp > ans {
					ans = tmp
				}
			}
		}
	}
	return ans
}
// @lc code=end

