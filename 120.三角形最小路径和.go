/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */
// package main

// func main() {

// }
// @lc code=start
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 0 {
		return 0
	}
	n := len(triangle[m-1])
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if j == 0 {
				dp[j] += triangle[i][j]
				continue
			}
			if j == len(triangle[i])-1 {
				dp[j] = dp[j-1] + triangle[i][j]
				continue
			}
			dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
		}
		// fmt.Println(dp)
	}
	ans := math.MaxInt32
	for j := 0; j < n; j++ {
		if ans > dp[j] {
			ans = dp[j]
		}
	}
	return ans
	// 2
	// 3 4
	// 6 5 7
	// 4 1 8 3
	// 一边前进一边修改数组就好了

	// dp := make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([]int, n)
	// }
	// dp[0][0] = triangle[0][0]
	// for i := 1; i < m; i++ {
	// 	for j := 0; j < len(triangle[i]); j++ {
	// 		if j == 0 {
	// 			dp[i][j] = dp[i-1][j] + triangle[i][j]
	// 			continue
	// 		}
	// 		if j == len(triangle[i])-1 {
	// 			dp[i][j] = dp[i-1][j-1] + triangle[i][j]
	// 			continue
	// 		}
	// 		a := dp[i-1][j-1]
	// 		b := dp[i-1][j]
	// 		dp[i][j] = min(a, b) + triangle[i][j]
	// 	}
	// }
	// fmt.Println(dp)
	// ans := dp[m-1][0]
	// for j := 0; j < len(triangle[m-1]); j++ {
	// 	if ans > dp[m-1][j] {
	// 		ans = dp[m-1][j]
	// 	}
	// }
	// return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
