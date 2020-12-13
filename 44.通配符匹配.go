/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */
package main

import "fmt"

// func main() {
// 	isMatch("aa", "a")
// 	isMatch("aa", "*")
// 	isMatch("cb", "?a")
// 	isMatch("adceb", "*a*b")
// 	isMatch("acdcb", "*a*c?b")
// }

// @lc code=start
func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// 如果使用了 p[j-1] 那么需要 dp[i-1][j]
				// 说明当前*可以匹配 s[i-1] ，那么 s[i-2] 之前的需要 dp[i-1][j]
				// 不使用 p[j-1] dp[i][j-1]
				// 说明当前的 * 是空字符串的意思那么需要前一段去匹配这个 s[i-1]

				// 1.看作匹配0个 s不动，p动， dp[i][j-1]
				// 2.看作匹配1个s动,p不动,dp[i-1][j]
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] && s[i-1] == p[j-1]
			}
		}
	}
	fmt.Println("ans", dp[m][n])
	fmt.Println("------next------")
	return dp[m][n]
}
func isMatchFail(s string, p string) bool {
	n := len(s)
	m := len(p)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, m)
	}
	// dp[0][0] = true
	// 如果前 s[:i] p[:j] 可以匹配，那么dp[i][j] = true
	// 1. p[j] = ? dp[i][j] = dp[i-1][j-1]

	// 2.p[j] = * // 可以后退 n 步 dp[i][j] = dp[i-x][j-1] x=[1,i]

	//3 p[j] = w check p[j] == s[i]
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = p[j] == '?' || p[j] == '*' || s[i] == p[j]
			} else if i == 0 {
				// 当且仅当 p[:j+1] 只有一个以下的 s[i] 其他都是 ? * 才可以 = true
				same := 0
				q := 0
				any := 0
				diff := 0
				for k := 0; k < j; k++ {
					if s[i] == p[k] {
						same++
					} else if p[k] == '?' {
						q++
					} else if p[k] == '*' {
						any++
					} else {
						diff++
					}
				}
				fmt.Println(j, any, q, same, diff)
				if p[j] == '?' {
					dp[i][j] = any == j
				} else if p[j] == '*' {
					dp[i][j] = any == j-1 && diff == 0 && ((q == 1 && same == 0) || (q == 0 && same == 1))
				} else {
					dp[i][j] = any == j && diff == 0 && s[i] == p[j]
				}
			} else if j == 0 {
				dp[i][j] = p[j] == '*'
			} else {
				if p[j] == '?' {
					dp[i][j] = dp[i-1][j-1]
				} else if p[j] == '*' {
					pass := false
					for x := 0; x <= i; x++ {
						if dp[i-x][j-1] {
							pass = true
							break
						}
					}
					dp[i][j] = pass
				} else {
					dp[i][j] = dp[i-1][j-1] && s[i] == p[j]
				}
			}
			// fmt.Println(s[:i+1], p[:j+1], dp[i][j])
		}
	}
	// fmt.Println("ans", dp[n-1][m-1])
	// fmt.Println("------next------")
	return dp[n-1][m-1]
}

// @lc code=end
