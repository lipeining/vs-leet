/*
 * @lc app=leetcode.cn id=576 lang=golang
 *
 * [576] 出界的路径数
 */

// @lc code=start
func findPaths(m int, n int, N int, i int, j int) int {
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
