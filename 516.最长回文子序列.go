/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 */
package main

import "fmt"

// func main() {
// 	longestPalindromeSubseq("bbbab")
// 	longestPalindromeSubseq("bbbaab")
// }

// @lc code=start
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 当我知道 dp[i][i] 的特殊情况时，
	// 应该使用 填写表格 的方式，
	// 判断 dp[i][j] 与旁边的单元格的关系 只和 +1， -1 或者多个有关。
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	ans := dp[0][n-1]
	fmt.Println("ans", ans)
	return ans
}

// @lc code=end
