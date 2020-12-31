/*
 * @lc app=leetcode.cn id=639 lang=golang
 *
 * [639] 解码方法 2
 */
package main

import "fmt"

// @lc code=start
func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	if s[0] == '*' {
		dp[1] = 9
	}
	if s[0] == '0' {
		dp[1] = 0
	}
	toMod := int(1e9 + 7)
	for i := 1; i < n; i++ {
		if s[i] == '*' {
			dp[i+1] = 9 * dp[i]
			if s[i-1] == '1' {
				dp[i+1] = (dp[i+1] + 9*dp[i-1]) % toMod
			} else if s[i-1] == '2' {
				dp[i+1] = (dp[i+1] + 6*dp[i-1]) % toMod
			} else if s[i-1] == '*' {
				dp[i+1] = (dp[i+1] + 15*dp[i-1]) % toMod
			}
		} else {
			dp[i+1] = 0
			if s[i] != '0' {
				dp[i+1] = dp[i]
			}
			if s[i-1] == '1' {
				dp[i+1] = (dp[i+1] + dp[i-1]) % toMod
			} else if s[i-1] == '2' && s[i] <= '6' {
				dp[i+1] = (dp[i+1] + dp[i-1]) % toMod
			} else if s[i-1] == '*' {
				last := 1
				if s[i] <= '6' {
					last = 2
				}
				dp[i+1] = (dp[i+1] + last*dp[i-1]) % toMod
			}
		}
	}
	fmt.Println("ans", dp)
	return dp[n]
}

// @lc code=end
