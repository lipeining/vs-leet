/*
 * @lc app=leetcode.cn id=1269 lang=golang
 *
 * [1269] 停在原地的方案数
 */
package main

// func main() {
// 	numWays(3, 2)
// 	numWays(2, 4)
// 	numWays(4, 2)
// }

// @lc code=start
func numWays(steps int, arrLen int) int {
	dp := make([][]int, steps+1)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	arrLen = min(arrLen, steps+1)
	for i := 0; i <= steps; i++ {
		dp[i] = make([]int, arrLen)
	}
	to := int(1e9 + 7)
	dp[0][0] = 1
	for step := 1; step <= steps; step++ {
		for j := 0; j < arrLen; j++ {
			left := 0
			right := 0
			if j > 0 {
				left = dp[step-1][j-1]
			}
			if j < arrLen-1 {
				right = dp[step-1][j+1]
			}
			dp[step][j] = (dp[step-1][j] + left + right) % to
		}
		// fmt.Println("step", step, dp[step])
	}
	// fmt.Println(dp)
	ans := dp[steps][0]
	// fmt.Println("anssssss", ans)
	return ans
}

// @lc code=end
