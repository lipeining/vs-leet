package main

import "fmt"

// func main() {
// 	maxProductPath([][]int{
// 		{-1, -2, -3},
// 		{-2, -3, -3},
// 		{-3, -3, -2},
// 	})
// 	maxProductPath([][]int{
// 		{1, -2, 1},
// 		{1, -2, 1},
// 		{3, -4, 1},
// 	})
// 	maxProductPath([][]int{
// 		{1, 3},
// 		{0, -4},
// 	})
// 	maxProductPath([][]int{
// 		{1, 4, 4, 0},
// 		{-2, 0, 0, 1},
// 		{1, -1, 1, 1},
// 	})
// }
func maxProductPath(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	dp := make([][][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([][]int, cols)
		for j := 0; j < cols; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	mod := int(1e9 + 7)
	// 0 为最小值，1为最大值
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 && j == 0 {
				dp[i][j][0] = grid[i][j]
				dp[i][j][1] = grid[i][j]
			} else if i == 0 {
				dp[i][j][0] = dp[i][j-1][0] * grid[i][j]
				dp[i][j][1] = dp[i][j-1][1] * grid[i][j]
			} else if j == 0 {
				dp[i][j][0] = dp[i-1][j][0] * grid[i][j]
				dp[i][j][1] = dp[i-1][j][1] * grid[i][j]
			} else {
				a := dp[i-1][j][0] * grid[i][j]
				b := dp[i-1][j][1] * grid[i][j]
				c := dp[i][j-1][0] * grid[i][j]
				d := dp[i][j-1][1] * grid[i][j]
				dp[i][j][0] = min(min(a, b), min(c, d))
				dp[i][j][1] = max(max(a, b), max(c, d))
			}
		}
	}
	ans := dp[rows-1][cols-1][1]
	fmt.Println("anssss", ans)
	if ans < 0 {
		return -1
	}
	return ans % mod
}
