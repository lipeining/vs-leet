/*
 * @lc app=leetcode.cn id=576 lang=golang
 *
 * [576] 出界的路径数
 */
package main

import "fmt"

// func main() {
// 	findPaths(2, 2, 2, 0, 0)
// 	findPaths(1, 3, 3, 0, 1)
// }

// @lc code=start
func findPaths(m int, n int, N int, i int, j int) int {
	// 在坐标 i,j 时移动 k 次可以出界的数量
	memo := make([][][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]int, N+1)
			for k := 0; k <= N; k++ {
				// memo[i][j][k] = -1
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				memo[i][j][1]++
			}
			if i == m-1 {
				memo[i][j][1]++
			}
			if j == 0 {
				memo[i][j][1]++
			}
			if j == n-1 {
				memo[i][j][1]++
			}
		}
	}
	fmt.Println("init", memo)
	mod := int(1e9 + 7)
	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	var dfs func(x, y, k int) int
	dfs = func(x, y, k int) int {
		if k <= 0 {
			return 0
		}
		if x < 0 || x == m || y < 0 || y >= n {
			fmt.Println("edge")
			return 0
		}
		if memo[x][y][k] > 0 {
			fmt.Println("memo", x, y, k)
			return memo[x][y][k]
		}
		// 1.移动到相邻的位置
		// 2.直接出界
		now := 0
		for _, dir := range dirs {
			ni := x + dir[0]
			nj := y + dir[1]
			now += dfs(ni, nj, k-1)
			now %= mod
		}
		memo[x][y][k] = now
		fmt.Println("dfs", x, y, k, now)
		return now
	}

	ans := dfs(i, j, N)
	fmt.Println("anssssss", ans)
	return ans
}
func findPathsDP(m int, n int, N int, i int, j int) int {
	if N == 0 {
		return 0
	}
	dp := make([][][]int, m+2)
	for i := 0; i < m+2; i++ {
		dp[i] = make([][]int, n+2)
		for j := 0; j < n+2; j++ {
			dp[i][j] = make([]int, N+1)
		}
	}
	// 别人的初始化思路
	// 初始化边界
	// 相当于加了两层保护的边界
	// https://leetcode-cn.com/problems/out-of-boundary-paths/solution/zhuang-tai-ji-du-shi-zhuang-tai-ji-by-christmas_wa/
	for i := 0; i <= m+1; i++ {
		dp[i][0][0] = 1
		dp[i][n+1][0] = 1
	}
	for i := 0; i <= n+1; i++ {
		dp[0][i][0] = 1
		dp[m+1][i][0] = 1
	}
	mod := int(1e9 + 7)
	ways := [4][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	// 上下左右四个方向 加一
	// 为什么是全部的 step=0 都为 1 呢？
	for step := 1; step <= N; step++ {
		for x := 1; x <= m; x++ {
			for y := 1; y <= n; y++ {
				sum := 0
				for _, way := range ways {
					dx, dy := way[0], way[1]
					nx, ny := x+dx, y+dy
					sum = (sum + dp[nx][ny][step-1]) % mod
				}
				dp[x][y][step] = sum
			}
		}
	}
	ans := 0
	for k := 1; k <= N; k++ {
		ans = (ans + dp[i+1][j+1][k]) % mod
	}
	return ans
}

// @lc code=end
