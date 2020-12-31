/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(A []int, B []int) int {
	m := len(A)
	n := len(B)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if ans < dp[i][j] {
					ans = dp[i][j]
				}
			}
		}
	}
	return ans
	// for i := m - 1; i >= 0; i-- {
	// 	for j := n - 1; j >= 0; j-- {
	// 		if A[i] == B[j] {
	// 			dp[i][j] = dp[i+1][j+1] + 1
	// 			if ans < dp[i][j] {
	// 				ans = dp[i][j]
	// 			}
	// 		}
	// 	}
	// }
	// return ans
}

// @lc code=end
