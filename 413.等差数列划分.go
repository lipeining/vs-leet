/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */

// @lc code=start
func numberOfArithmeticSlices(A []int) int {
	dp := 0
	sum := 0
	for i := 2; i < len(A); i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp = dp + 1
			sum += dp
		} else {
			dp = 0
		}
	}
	return sum
	// length := len(A)
	// dp := make([]int, length)
	// sum := 0
	// for i:=2;i<length;i++ {
	// 	if A[i]-A[i-1] == A[i-1]-A[i-2] {
	// 		dp[i] = dp[i-1]+1
	// 	} else {
	// 		dp[i] = 0
	// 	}
	// 	sum += dp[i]
	// }
	// return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
