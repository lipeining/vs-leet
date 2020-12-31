/*
 * @lc app=leetcode.cn id=688 lang=golang
 *
 * [688] “马”在棋盘上的概率
 */

// @lc code=start
func knightProbability(N int, K int, r int, c int) float64 {
	dp := make([][][]float64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([][]float64, N)
		for j := 0; j < N; j++ {
			dp[i][j] = make([]float64, K+1)
		}
	}
	// 根据题目我们可以知道 (dr, dc)(dr,dc) 的可能数据对是
	// (2, 1),(2,1), (2, -1),(2,−1), (-2, 1),
	// (−2,1), (-2, -1),(−2,−1), (1, 2),(1,2),
	//  (1, -2),(1,−2), (-1, 2),(−1,2), (-1, -2)(−1,−2)。
	// 我们将使用二维的 dp 和 dp2 来存储我们的数据，而不是使用三维数组 f。
	// dp2 代表 f[][][steps]，dp 代表 f[][][steps-1]。
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
	dp[r][c][0] = 1.0
	// 初始化时，只有 r,c 为 1， 之后不断地跳动，直到最后一步，ans 为棋盘上的概率总和
	for step := 0; step < K; step++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				for _, way := range ways {
					dx, dy := way[0], way[1]
					x, y := i+dx, j+dy
					if x >= 0 && x < N && y >= 0 && y < N {
						dp[x][y][step+1] += dp[i][j][step] / 8.0
					}
				}
			}
		}
	}
	// fmt.Println(dp)
	var ans float64
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ans += dp[i][j][K]
		}
	}
	return ans
	// 旧的写法
	// fmt.Println(ways)
	// return 0.0
	// dp[r][c][0] = 1.0
	// 为什么是全部的 step=0 都为 1 呢？
	// for step := 0; step <= K; step++ {
	// 	for i := 0; i < N; i++ {
	// 		for j := 0; j < N; j++ {
	// 			if step == 0 {
	// 				dp[i][j][step] = 1.0
	// 			} else {
	// 				sum := 0.0
	// 				for _, way := range ways {
	// 					dx, dy := way[0], way[1]
	// 					x, y := i+dx, j+dy
	// 					if x >= 0 && x < N && y >= 0 && y < N {
	// 						sum += dp[x][y][step-1]
	// 					}
	// 				}
	// 				dp[i][j][step] = sum / 8.0
	// 			}
	// 		}
	// 	}
	// }
	// return dp[r][c][K]
}

// @lc code=end
