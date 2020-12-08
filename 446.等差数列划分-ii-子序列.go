/*
 * @lc app=leetcode.cn id=446 lang=golang
 *
 * [446] 等差数列划分 II - 子序列
 */
package main

import "fmt"

// func main() {
// 	numberOfArithmeticSlices([]int{2, 4, 6, 8, 10})
// }

// @lc code=start
func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	dp := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
	}
	// A[i] 以公差 j 结束的数量
	ans := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := A[i] - A[j]
			dp[i][diff] += dp[j][diff] + 1
			// 长度为 3，4，5 之后有所不同
			ans += dp[j][diff]
			// fmt.Println("i, j, diff, dp[i]", i, j, diff, dp[i])
		}
	}
	// for i := 1; i < n; i++ {
	// 	for j := 0; j < i; j++ {
	// 		diff := A[i] - A[j]
	// 		if dp[j][diff] != 0 {
	// 			dp[i][diff] = dp[j][diff] + 1
	// 		} else {
	// 			dp[i][diff] = 2
	// 		}
	// 		if dp[i][diff] >= 3 {
	// 			ans += dp[i][diff] - 3 + 1
	// 		}
	// 		// fmt.Println("i, j, diff, dp[i]", i, j, diff, dp[i])
	// 	}
	// }
	fmt.Println("ans", ans)
	return ans
}

// @lc code=end
