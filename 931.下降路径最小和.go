import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=931 lang=golang
 *
 * [931] 下降路径最小和
 */

// @lc code=start
func minFallingPathSum(A [][]int) int {
	length := len(A)
	for i := 1; i < length; i++ {
		for j := 0; j < length; j++ {
			if j == 0 {
				A[i][j] = min(A[i-1][j], A[i-1][j+1]) + A[i][j]
			} else if j == length-1 {
				A[i][j] = min(A[i-1][j], A[i-1][j-1]) + A[i][j]
			} else {
				A[i][j] = min3(A[i-1][j-1], A[i-1][j], A[i-1][j+1]) + A[i][j]
			}
		}
	}
	// fmt.Println(A)
	ans := math.MaxInt32
	for j := 0; j < length; j++ {
		ans = min(A[length-1][j], ans)
	}
	return ans

	// length := len(A)
	// dp := make([][]int, length)
	// for i := 0; i < length; i++ {
	// 	dp[i] = make([]int, length)
	// }
	// for j := 0; j < length; j++ {
	// 	dp[0][j] = A[0][j]
	// }
	// for i := 1; i < length; i++ {
	// 	for j := 0; j < length; j++ {
	// 		if j == 0 {
	// 			dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + A[i][j]
	// 		} else if j == length-1 {
	// 			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + A[i][j]
	// 		} else {
	// 			dp[i][j] = min3(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) + A[i][j]
	// 		}
	// 	}
	// }
	// // fmt.Println(dp)
	// ans := math.MaxInt32
	// for j := 0; j < length; j++ {
	// 	ans = min(dp[length-1][j], ans)
	// }
	// return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func min3(a, b, c int) int {
	return min(a, min(b, c))
}

// @lc code=end
