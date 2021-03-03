/*
 * @lc app=leetcode.cn id=1289 lang=golang
 *
 * [1289] 下降路径最小和  II
 */
package main

// import "fmt"

// func main() {
// 	minFallingPathSum([][]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 9},
// 	})
// }

// @lc code=start
func minFallingPathSum(arr [][]int) int {
	n := len(arr)
	dp := make([][]int, n)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	ans := 100 * n
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = ans
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = arr[i][j]
			} else {
				for k := 0; k < n; k++ {
					if k != j {
						dp[i][j] = min(dp[i][j], dp[i-1][k]+arr[i][j])
					}
				}
			}
			if i == n-1 {
				ans = min(ans, dp[i][j])
			}
		}
		// fmt.Println(dp[i])
	}
	return ans
}

// @lc code=end
