/*
 * @lc app=leetcode.cn id=778 lang=golang
 *
 * [778] 水位上升的泳池中游泳
 */
package main

// func main() {
// 	swimInWater([][]int{{0, 2}, {1, 3}})
// 	swimInWater([][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}})
// }

// @lc code=start
func swimInWater(grid [][]int) int {
	n := len(grid)
	l := 0
	r := n * n
	toId := func(i, j int) int {
		return i*n + j
	}
	toIJ := func(id int) (int, int) {
		return id / n, id % n
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	check := func(t int) bool {
		//是否存在从 0,0 到 n,n 的路径
		queue := make([]int, 0)
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, n)
		}
		if grid[0][0] <= t {
			queue = append(queue, 0)
			visited[0][0] = true
		}
		for len(queue) != 0 {
			size := len(queue)
			// fmt.Println("queue", queue)
			for k := 0; k < size; k++ {
				i, j := toIJ(queue[k])
				if i == n-1 && j == n-1 {
					return true
				}
				for _, dir := range dirs {
					x, y := i+dir[0], j+dir[1]
					if x < 0 || x == n || y < 0 || y == n || visited[x][y] {
						continue
					}
					visited[x][y] = true
					if grid[x][y] <= t {
						queue = append(queue, toId(x, y))
					}
				}
			}
			queue = queue[size:]
		}
		return false
	}
	for l < r {
		mid := (l + r) >> 1
		// fmt.Println("mid", mid)
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// fmt.Println("ansssssssssss", l)
	return l
}

// @lc code=end
