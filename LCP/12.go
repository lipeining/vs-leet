package main

import "fmt"

func main() {
	minTime([]int{1, 2, 3, 3}, 2)
	minTime([]int{999, 999, 999}, 4)
}
func minTime(time []int, m int) int {
	n := len(time)
	if m >= n {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// 前 i 天完成 j 道题目时，花费的最大一天时间。
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + time[i]
	}
	area := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		area[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		area[i][i] = time[i-1]
		for j := i + 1; j <= n; j++ {
			area[i][j] = max(area[i][j-1], time[j-1])
		}
	}
	fmt.Println(pre)
	fmt.Println(area)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k < j; k++ {
				dp[i][j] = max(dp[i-1][k], pre[j]-pre[k]-area[k][j])
			}
		}
	}
	fmt.Println(dp, dp[m][n])
	return dp[m][n]
}
