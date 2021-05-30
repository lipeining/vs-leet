/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 */

// @lc code=start
func orangesRotting(grid [][]int) int {
	// bfs 广度优先
	// var bfs func(grid [][]int, i int, j int)
	// bfs = func(grid [][]int, i int, j int) {
	//     if i < 0 || i >=m || j < 0 || j >= n {
	//         return
	//     }
	//     if grid[i][j] != 2 {
	//         return
	//     }
	//     bfs(grid, i-1, j)
	//     bfs(grid, i+1, j)
	//     bfs(grid, i, j-1)
	//     bfs(grid, i, j+1)
	// }
	m := len(grid)
	n := len(grid[0])
	ans := 0
	queue := make([][]int, 0)
	normal := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				normal++
			}
		}
	}
	if normal == 0 {
		return 0
	}
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) != 0 {
		length := len(queue)
		for k := 0; k < length; k++ {
			tmp := queue[k]
			i := tmp[0]
			j := tmp[1]
			for _, dir := range dirs {
				x := i + dir[0]
				y := j + dir[1]
				if x >= 0 && x < m && y >= 0 && y < n {
					if grid[x][y] == 1 {
						grid[x][y] = 2
						normal--
						queue = append(queue, []int{x, y})
					}
				}
			}
		}
		ans++
		queue = queue[length:]
	}
	if normal != 0 {
		return -1
	}
	return ans - 1
}

// @lc code=end

