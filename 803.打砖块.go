/*
 * @lc app=leetcode.cn id=803 lang=golang
 *
 * [803] 打砖块
 */
//  [[1],[1],[1],[1],[1]]
//  [[3,0],[4,0],[1,0],[2,0],[0,0]]
package main

// func main() {
// 	hitBricks([][]int{{1}, {1}, {1}, {1}, {1}}, [][]int{{3, 0}, {4, 0}, {1, 0}, {2, 0}, {0, 0}})
// }

// @lc code=start
func hitBricks(grid [][]int, hits [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	ans := make([]int, len(hits))
	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				total++
			}
		}
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i == m || j < 0 || j == n {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		if grid[i][j] == 2 {
			return
		}
		grid[i][j] = 2
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	var count func() int
	count = func() int {
		for j := 0; j < n; j++ {
			if grid[0][j] == 1 {
				dfs(0, j)
			}
		}
		now := 0
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 2 {
					grid[i][j] = 1
					now++
				} else {
					grid[i][j] = 0
				}
			}
		}
		remove := total - now
		total = now
		return remove
	}
	for index, hit := range hits {
		target := grid[hit[0]][hit[1]]
		if target == 1 {
			grid[hit[0]][hit[1]] = 0
			total--
			ans[index] = count()
		}
	}
	// fmt.Println(ans)
	return ans
}

// @lc code=end
