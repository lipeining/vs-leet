/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */
package main

import "fmt"

// func main() {
// 	numDistinct("rabbbit", "rabbit")
// 	numDistinct("babgbag", "bag")
// }

// @lc code=start
func numDistinct(s string, t string) int {
	n := len(s)
	m := len(t)
	if m > n {
		return 0
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	// dp[i][j] s[:i]个字符能够组成 t[:j] 的方案数
	// 是否需要保证 i 为结尾呢？
	// 是否使用 s[i] 根据这个防线进行编写 不使用的话，为 dp[i-1][j]
	dp[0][0] = 1
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if j == 0 {
				dp[i][j] = 1
			}
			if i > 0 && j > 0 {
				if s[i-1] == t[j-1] {
					dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j]
				}
				// fmt.Println("i, j, dp[i][j]", i, j, dp[i][j])
			}
		}
	}
	fmt.Println("ans", dp[n][m])
	fmt.Println("------------------------")
	return dp[n][m]
}

// @lc code=end
