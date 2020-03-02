import "math"

/*
 * @lc app=leetcode.cn id=1039 lang=golang
 *
 * [1039] 多边形三角剖分的最低得分
 */

// @lc code=start
func minScoreTriangulation(A []int) int {
	// 有一个循环的操作，如何处理呢
	// dp[i][j] = min(dp[i][j], dp[start][k] + (A[start]*A[k]*A[k+1])+ dp[k+1][j])
	length := len(A)
	if length < 3 {
		return 0
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		for j := 0; j < length; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	// 初始化数据
	for i := 0; i < length; i++ {
		dp[i][(i+1)%length] = 0
	}
	// dp[i][j] = min{dp[i][k] + dp[k][j] + A[i][k][j]}
	// for all k in range (i, j)
	for l := 2; l < length; l++ {
		for i := 0; i < length-l; i++ {
			j := (i + l) % length
			for k := (i + 1) % l; k != j; k = (k + 1) % length {
				dp[i][j] = min(dp[i][j], dp[i][k]+A[i]*A[k]*A[j]+dp[k][j])
			}
		}
	}
	return dp[0][length-1]
	// // 初始化数据
	// for i := 0; i < length; i++ {
	// 	j := i + 2
	// 	dp[i][(i+1)%length] = 0
	// 	if j < length {
	// 		dp[i][j] = A[i] * A[i+1] * A[i+2]
	// 	}
	// }
	// for l := 2; l < length; l++ {
	// 	for i := 0; i < length-l; i++ {
	// 		j := i + l
	// 		for k := i + 2; k < j; k++ {
	// 			dp[i][j] = min(dp[i][j], dp[i][k]+A[i]*A[k]*A[k+1]+dp[k+1][j])
	// 		}
	// 	}
	// }
	// return dp[0][length-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
