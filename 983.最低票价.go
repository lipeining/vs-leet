import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=983 lang=golang
 *
 * [983] 最低票价
 */

// @lc code=start
func mincostTickets(days []int, costs []int) int {
	length := len(days)
	target := days[length-1]
	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = -1
	}
	for i := 0; i < len(days); i++ {
		dp[days[i]] = math.MaxInt32
	}
	jMap := map[int]int{0: 1, 1: 7, 2: 30}
	// 和硬币的类似，不过，提前购买是如何处理
	for i := 1; i <= target; i++ {
		if dp[i] != -1 {
			for j := 0; j < len(costs); j++ {
				cost := costs[j]
				day, _ := jMap[j]
				if day <= i {
					dp[i] = min(dp[i], dp[i-day]+cost)
				} else {
					// 可以提前购买
					dp[i] = min(dp[i], cost)
				}
			}
		} else {
			dp[i] = dp[i-1]
		}
	}
	// fmt.Println(dp)
	return dp[target]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
