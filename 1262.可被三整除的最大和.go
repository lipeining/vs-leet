/*
 * @lc app=leetcode.cn id=1262 lang=golang
 *
 * [1262] 可被三整除的最大和
 */
package main

import (
	"fmt"
	"math"
)

// func main() {
// 	maxSumDivThree([]int{3, 6, 5, 1, 8})
// 	maxSumDivThree([]int{4})
// 	maxSumDivThree([]int{1, 2, 3, 4, 4})
// }

// @lc code=start
func maxSumDivThree(nums []int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 3)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp[0][0] = 0
	dp[0][1] = math.MinInt32
	dp[0][2] = math.MinInt32
	for i := 1; i <= n; i++ {
		if nums[i-1]%3 == 0 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][0]+nums[i-1])
			dp[i][1] = max(dp[i-1][1], dp[i-1][1]+nums[i-1])
			dp[i][2] = max(dp[i-1][2], dp[i-1][2]+nums[i-1])
		} else if nums[i-1]%3 == 1 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][2]+nums[i-1])
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]+nums[i-1])
			dp[i][2] = max(dp[i-1][2], dp[i-1][1]+nums[i-1])
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+nums[i-1])
			dp[i][1] = max(dp[i-1][1], dp[i-1][2]+nums[i-1])
			dp[i][2] = max(dp[i-1][2], dp[i-1][0]+nums[i-1])
		}
	}
	// fmt.Println("anssss", dp[n][0])
	return dp[n][0]
}
func maxSumDivThreeTime(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)
	}
	// dp[i][1] 以 i 结尾的 余数为 1 的最大和
	//
	// max := func(a, b int) int {
	// 	if a > b {
	// 		return a
	// 	}
	// 	return b
	// }
	ans := 0
	// 需要使用最大值记录上一个节点，加快搜索速度
	// 优化为 nlogn
	for i := 0; i < n; i++ {
		right := nums[i] % 3
		dp[i][right] = nums[i]
		if i == 0 {
			continue
		}
		// timeout code 超时代码
		// fmt.Println("I after ", i, dp[i])
		// fmt.Println("I before", i, dp[i])
		// for j := 0; j < i; j++ {
		// 	// for mod := 0; mod <= 2; mod++ {
		// 	// 	dp[i][(right+mod)%3] = max(dp[i][(right+mod)%3], nums[i]+dp[j][mod])
		// 	// }
		// 	// fmt.Println("j", j, dp[j])
		// 	// newZero := max(dp[i][0]+dp[j][0], max(dp[i][1]+dp[j][2], dp[i][2]+dp[j][1]))
		// 	// newOne := max(dp[i][1]+dp[j][0], max(dp[i][0]+dp[j][1], dp[i][2]+dp[j][2]))
		// 	// newTwo := max(dp[i][2]+dp[j][0], max(dp[i][1]+dp[j][1], dp[i][2]+dp[j][0]))
		// 	// dp[i][0] = newZero
		// 	// dp[i][1] = newOne
		// 	// dp[i][2] = newTwo
		// 	if right == 0 {
		// 		dp[i][0] = max(dp[i][0], nums[i]+dp[j][0])
		// 		if dp[j][1] != 0 {
		// 			dp[i][1] = max(dp[i][1], nums[i]+dp[j][1])
		// 		}
		// 		if dp[j][2] != 0 {
		// 			dp[i][2] = max(dp[i][2], nums[i]+dp[j][2])
		// 		}
		// 	} else if right == 1 {
		// 		if dp[j][2] != 0 {
		// 			dp[i][0] = max(dp[i][0], nums[i]+dp[j][2])
		// 		}
		// 		dp[i][1] = max(dp[i][1], nums[i]+dp[j][0])
		// 		if dp[j][1] != 0 {
		// 			dp[i][2] = max(dp[i][2], nums[i]+dp[j][1])
		// 		}

		// 	} else {
		// 		if dp[j][1] != 0 {
		// 			dp[i][0] = max(dp[i][0], nums[i]+dp[j][1])
		// 		}
		// 		if dp[j][2] != 0 {
		// 			dp[i][1] = max(dp[i][1], nums[i]+dp[j][2])
		// 		}
		// 		dp[i][2] = max(dp[i][2], nums[i]+dp[j][0])
		// 	}
		// }
		// ans = max(ans, dp[i][0])
		// fmt.Println("I after ", i, dp[i])
	}
	fmt.Println("ansssss", ans)
	return ans
}

// @lc code=end
