/*
 * @lc app=leetcode.cn id=935 lang=golang
 *
 * [935] 骑士拨号器
 */

// @lc code=start
func knightDialer(N int) int {
	ways := [8][2]int{
		{2, 1},
		{2, -1},
		{-2, 1},
		{-2, -1},
		{1, 2},
		{1, -2},
		{-1, 2},
		{-1, -2},
	}
	dp := make([][][]int, 4)
	for i := 0; i < 4; i++ {
		dp[i] = make([][]int, 3)
		for j := 0; j < 3; j++ {
			dp[i][j] = make([]int, N)
		}
	}
	mod := int(1e9 + 7)
	// 为什么是 N，因为只能跳 N - 1 步
	for step := 0; step < N; step++ {
		for i := 0; i < 4; i++ {
			for j := 0; j < 3; j++ {
				checkRes := checkNumber(i, j)
				if !checkRes {
					continue
				} else {
					if step == 0 {
						dp[i][j][step] = 1
					} else {
						sum := 0
						for _, way := range ways {
							dx, dy := way[0], way[1]
							x, y := i+dx, j+dy
							if checkNumber(x, y) {
								sum = (sum + dp[x][y][step-1]) % mod
							}
						}
						dp[i][j][step] = sum
					}
				}
			}
		}
	}
	// fmt.Println(dp)
	ans := 0
	// 最后这里也需要 ans % mod
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			checkRes := checkNumber(i, j)
			if !checkRes {
				continue
			} else {
				ans = (ans + dp[i][j][N-1]) % mod
			}
		}
	}
	return ans
}
func checkNumber(x, y int) bool {
	if x == 3 && y == 1 {
		return true
	}
	return x >= 0 && x < 3 && y >= 0 && y < 3
}

// @lc code=end
