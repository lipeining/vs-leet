/*
 * @lc app=leetcode.cn id=1278 lang=golang
 *
 * [1278] 分割回文串 III
 */
package main

import (
	"math"
)

// func main() {
// 	palindromePartition("abc", 2)
// 	palindromePartition("aabbc", 3)
// 	palindromePartition("leetcode", 8)
// }

// @lc code=start
func palindromePartition(s string, k int) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
		for x := 0; x <= k; x++ {
			dp[i][x] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// need := func(left, right int) int {
	// 	cnt := 0
	// 	// fmt.Println("need j, i", left, right, s[left:right])
	// 	for left < right {
	// 		if s[left] != s[right] {
	// 			cnt++
	// 		}
	// 		left++
	// 		right--
	// 	}
	// 	return cnt
	// }
	// 他人答案
	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
	}
	for span := 2; span <= n; span++ {
		for i := 0; i <= n-span; i++ {
			j := i + span - 1
			equal := 0
			if s[i] != s[j] {
				equal = 1
			}
			cost[i][j] = cost[i+1][j-1] + equal
		}
	}
	for i := 1; i <= n; i++ {
		for x := 1; x <= min(k, i); x++ {
			if x == 1 {
				// dp[i][x] = need(0, i-1)
				dp[i][x] = cost[0][i-1]
			} else {
				for i0 := x - 1; i0 < i; i0++ {
					// dp[i][x] = min(dp[i][x], dp[i0][x-1]+need(i0, i-1))
					dp[i][x] = min(dp[i][x], dp[i0][x-1]+cost[i0][i-1])
				}
			}
		}
	}
	// fmt.Println(dp)
	// fmt.Println(dp[n][k])
	return dp[n][k]
}

// @lc code=end
