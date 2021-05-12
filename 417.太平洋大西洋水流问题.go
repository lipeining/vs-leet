/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 */
package main

// import "fmt"

// func main() {
// 	pacificAtlantic([][]int{
// 		{1, 2, 2, 3, 5},
// 		{3, 2, 3, 4, 4},
// 		{2, 4, 5, 3, 1},
// 		{6, 7, 1, 4, 5},
// 		{5, 1, 1, 2, 4},
// 	})
// }

// @lc code=start
func pacificAtlantic(matrix [][]int) [][]int {
	m := len(matrix)
	ans := make([][]int, 0)
	if m == 0 {
		return ans
	}
	n := len(matrix[0])
	if n == 0 {
		return ans
	}
	lt := make([][]bool, m)
	rb := make([][]bool, m)
	for i := 0; i < m; i++ {
		lt[i] = make([]bool, n)
		rb[i] = make([]bool, n)
	}
	illegal := func(x, y int) bool {
		return x < 0 || x >= m || y < 0 || y >= n
	}
	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	var flowlt func(x, y int)
	flowlt = func(x, y int) {
		if lt[x][y] {
			return
		}
		lt[x][y] = true
		for _, dir := range dirs {
			nx := x + dir[0]
			ny := y + dir[1]
			if !illegal(nx, ny) {
				if matrix[nx][ny] >= matrix[x][y] {
					flowlt(nx, ny)
				}
			}
		}
	}
	var flowrb func(x, y int)
	flowrb = func(x, y int) {
		if rb[x][y] {
			return
		}
		rb[x][y] = true
		for _, dir := range dirs {
			nx := x + dir[0]
			ny := y + dir[1]
			if !illegal(nx, ny) {
				if matrix[nx][ny] >= matrix[x][y] {
					flowrb(nx, ny)
				}
			}
		}
	}
	trylt := func() {
		for i := 0; i < m; i++ {
			flowlt(i, 0)
		}
		for j := 0; j < n; j++ {
			flowlt(0, j)
		}
	}
	tryrb := func() {
		for i := 0; i < m; i++ {
			flowrb(i, n-1)
		}
		for j := 0; j < n; j++ {
			flowrb(m-1, j)
		}
	}
	trylt()
	// fmt.Println(lt)
	tryrb()
	// fmt.Println(rb)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if lt[i][j] && rb[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	// fmt.Println(ans)
	return ans
}

// @lc code=end
