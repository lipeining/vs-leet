/*
 * @lc app=leetcode.cn id=1575 lang=golang
 *
 * [1575] 统计所有可行路径
 */
package main

import "fmt"

// func main() {
// 	countRoutes([]int{2, 3, 6, 8, 4}, 1, 3, 5)
// 	countRoutes([]int{4, 3, 1}, 1, 0, 6)
// 	countRoutes([]int{5, 2, 1}, 0, 2, 3)
// 	countRoutes([]int{2, 1, 5}, 0, 0, 3)
// 	countRoutes([]int{1, 2, 3}, 0, 2, 40)
// }

// @lc code=start
func countRoutes(locations []int, start int, finish int, fuel int) int {
	n := len(locations)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, fuel+1)
	}
	mod := int(1e9 + 7)
	ans := 0
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	// fuel 为 0 时，固定为 0
	// 从 i 到 finish 的 need > fuel，固定为 0
	for i := 0; i <= fuel; i++ {
		dp[finish][i] = 1
	}
	dp[finish][0] = 1
	for f := 0; f <= fuel; f++ {
		for i := 0; i < n; i++ {
			// if i != finish {
			// 	need := abs(localtions[i] - locations[finish])
			// 	if f >= need {
			// 		dp[i][f] += dp[finish][f-need]
			// 	}
			// }
			for j := 0; j < n; j++ {
				if i != j {
					need := abs(locations[i] - locations[j])
					if f >= need {
						dp[i][f] += dp[j][f-need]
						dp[i][f] %= mod
					}
				}
			}
		}
	}
	// fmt.Println(dp)
	ans = dp[start][fuel]
	fmt.Println("anssssss", ans)
	return ans % mod
}

// @lc code=end
