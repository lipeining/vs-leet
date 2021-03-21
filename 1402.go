package main

import (
	"fmt"
	"math"
	"sort"
)

// func main() {
// 	maxSatisfaction([]int{-1, -8, 0, 5, -9})
// 	maxSatisfaction([]int{4, 3, 2})
// 	maxSatisfaction([]int{-1, -4, -5})
// 	maxSatisfaction([]int{-2, 5, -1, 0, 3, -3})
// }
func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	// 最后面的得分最大
	// 计算 是否需要做第 i 道菜？
	// 到 i 为止，一共做了多少道菜。
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(satisfaction)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			dp[i][j] = math.MinInt32
		}
	}
	// dp[1][1]=satisfaction[0]
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+j*satisfaction[i-1])
			ans = max(ans, dp[i][j])
		}
		fmt.Println("i=", i, dp[i])
	}
	fmt.Println("anssssssssssss", ans)
	return ans
}
